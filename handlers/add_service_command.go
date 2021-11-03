package handlers

import (
	"errors"
	"evil-meow/owl-of-athena/config"
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/k8s"
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"gopkg.in/yaml.v2"
)

// handleHelloCommand will take care of /add_service submissions
func HandleAddServiceCommand(command slack.SlashCommand, client *slack.Client) error {
	serviceName := &command.Text
	username := &command.UserName
	channelID := &command.ChannelID

	log.Printf("Adding service %s...", *serviceName)

	sendMessage(client, channelID, serviceName, fmt.Sprintf("%s requested adding the service %s", *username, *serviceName))

	if github_api.IsGithubRepoCreated(*serviceName) {
		sendMessage(client, channelID, serviceName, fmt.Sprintf("Repo %s exists", *serviceName))
	} else {
		sendMessage(client, channelID, serviceName, fmt.Sprintf("Repo http://github.com/evil-meow/%s does not exist. Please, specify an existing repo.", *serviceName))
		return errors.New("base repo not found")
	}

	config, err := readConfigFile(serviceName)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not find owl.yml at the root of the repo. Please, create it in order to add the service.")
		return err
	}

	infraRepoName := *serviceName + "-infra"

	if github_api.IsGithubRepoCreated(infraRepoName) {
		sendMessage(client, channelID, serviceName, fmt.Sprintf("Infra repo %s already exists", infraRepoName))
	} else {
		sendMessage(client, channelID, serviceName, fmt.Sprintf("Infra repo %s does not exist. Creating.", infraRepoName))
		github_api.CreateGitubRepo(&infraRepoName)
	}

	err = commitReadme(&infraRepoName)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit README to the infra repo")
		return err
	}

	err = commitK8sDescriptors(&infraRepoName, config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit k8s descriptors to the infra repo")
		return err
	}

	return nil
}

func sendMessage(client *slack.Client, channelID *string, serviceName *string, text string) {
	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Service",
			Value: *serviceName,
		},
	}

	attachment.Text = text
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(*channelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Printf("failed to post message: %v", err)
	}
}

func readConfigFile(serviceName *string) (*config.Config, error) {
	configFileUrl := fmt.Sprintf("https://raw.githubusercontent.com/evil-meow/%s/main/owl.yml", *serviceName)
	configFile, err := github_api.ReadFile(&configFileUrl)
	if err != nil {
		log.Printf("Could not find config file at: %s", configFileUrl)
		return nil, errors.New("owl.yml file not found")
	}

	log.Printf("owl.yml found. Contents:\n%s", configFile)

	conf := config.Config{}

	yaml.Unmarshal([]byte(configFile), &conf)
	return &conf, nil
}

func commitReadme(repoName *string) error {
	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "README.md",
				Content:  "Repo created by owl-of-athena",
			},
		},
	}

	err := github_api.CommitFilesToMain(repoName, files)

	return err
}

func commitK8sDescriptors(repoName *string, config *config.Config) error {
	deploymentYaml, err := k8s.BuildDeploymentYaml(config)
	if err != nil {
		return err
	}

	namespaceYaml, err := k8s.BuildNamespaceYaml(config)
	if err != nil {
		return err
	}

	kustomizeYaml, err := k8s.BuildKustomizeYaml(config)
	if err != nil {
		return err
	}

	kustomizeProdYaml, err := k8s.BuildKustomizeProdYaml(config)
	if err != nil {
		return err
	}

	secretsProdYaml, err := k8s.BuildSecretsProdYaml(config)
	if err != nil {
		return err
	}

	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "base/deployment.yaml",
				Content:  deploymentYaml,
			},
			{
				FilePath: "base/namespace.yaml",
				Content:  namespaceYaml,
			},
			{
				FilePath: "base/kustomize.yaml",
				Content:  kustomizeYaml,
			},
			{
				FilePath: "overlays/production/kustomize.yaml",
				Content:  kustomizeProdYaml,
			},
			{
				FilePath: "overlays/production/deployment-secrets.yaml",
				Content:  secretsProdYaml,
			},
		},
	}

	err = github_api.CommitFilesToMain(repoName, files)

	return err
}
