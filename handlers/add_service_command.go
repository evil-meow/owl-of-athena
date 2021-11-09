package handlers

import (
	"errors"
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/k8s_api"
	"evil-meow/owl-of-athena/service_config"
	"evil-meow/owl-of-athena/templates/argocd"
	"evil-meow/owl-of-athena/templates/k8s"
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

	err = commitK8sDescriptors(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit k8s descriptors to the infra repo")
		return err
	}

	err = commitArgocd(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit argocd descriptor")
		return err
	}

	err = applyArgocd(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not apply argocd descriptor")
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
		log.Printf("Failed to post message: %v", err)
	}
}

func readConfigFile(serviceName *string) (*service_config.ServiceConfig, error) {
	configFileUrl := fmt.Sprintf("https://raw.githubusercontent.com/evil-meow/%s/main/owl.yml", *serviceName)
	configFile, err := github_api.ReadFile(&configFileUrl)
	if err != nil {
		log.Printf("Could not find config file at: %s", configFileUrl)
		return nil, errors.New("owl.yml file not found")
	}

	conf := service_config.ServiceConfig{}

	yaml.Unmarshal([]byte(configFile), &conf)
	conf.RepoName = *serviceName + "-infra"

	return &conf, nil
}

func commitReadme(repoName *string) error {
	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "README.md",
				Content:  "Repo created by owl-of-athena\n\nYou have a kustomize folder to use with kustomize and an argoCD application CRD.",
			},
		},
	}

	err := github_api.CommitFilesToMain(repoName, files)

	return err
}

func commitK8sDescriptors(config *service_config.ServiceConfig) error {
	deploymentYaml, err := k8s.BuildDeploymentYaml(config)
	if err != nil {
		return err
	}

	namespaceYaml, err := k8s.BuildNamespaceYaml(config)
	if err != nil {
		return err
	}

	serviceYaml, err := k8s.BuildServiceYaml(config)
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

	deploymentSecretsProdYaml, err := k8s.BuildDeploymentSecretsProdYaml(config)
	if err != nil {
		return err
	}

	secretsProdYaml, err := k8s.BuildSecretsProdYaml(config)
	if err != nil {
		return err
	}

	certificateProdYaml, err := k8s.BuildCertificateProdYaml(config)
	if err != nil {
		return err
	}

	gatewayProdYaml, err := k8s.BuildGatewayProdYaml(config)
	if err != nil {
		return err
	}

	labelProdYaml, err := k8s.BuildLabelProdYaml(config)
	if err != nil {
		return err
	}

	virtualServiceProdYaml, err := k8s.BuildVirtualServiceProdYaml(config)
	if err != nil {
		return err
	}

	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "kustomize/base/deployment.yaml",
				Content:  deploymentYaml,
			},
			{
				FilePath: "kustomize/base/namespace.yaml",
				Content:  namespaceYaml,
			},
			{
				FilePath: "kustomize/base/kustomization.yaml",
				Content:  kustomizeYaml,
			},
			{
				FilePath: "kustomize/base/service.yaml",
				Content:  serviceYaml,
			},
			{
				FilePath: "kustomize/overlays/production/kustomization.yaml",
				Content:  kustomizeProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/certificate.yaml",
				Content:  certificateProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/secrets.yaml",
				Content:  secretsProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/deployment-secrets.yaml",
				Content:  deploymentSecretsProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/gateway.yaml",
				Content:  gatewayProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/deployment-label.yaml",
				Content:  labelProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/virtual-service.yaml",
				Content:  virtualServiceProdYaml,
			},
		},
	}

	err = github_api.CommitFilesToMain(&config.RepoName, files)

	return err
}

func commitArgocd(config *service_config.ServiceConfig) error {
	argoYaml, err := argocd.BuildApplicationYaml(config)
	if err != nil {
		return err
	}

	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "argocd.yaml",
				Content:  argoYaml,
			},
		},
	}

	err = github_api.CommitFilesToMain(&config.RepoName, files)

	return err
}

func applyArgocd(config *service_config.ServiceConfig) error {
	err := k8s_api.Apply("https://github.com/evil-meow/" + config.RepoName + "/argocd.yaml")
	return err
}
