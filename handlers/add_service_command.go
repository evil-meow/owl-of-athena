package handlers

import (
	"errors"
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/k8s_api"
	"evil-meow/owl-of-athena/operations"
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
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

	config, err := operations.ReadConfigFile(serviceName)
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

	err = operations.CommitK8sDescriptors(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit k8s descriptors to the infra repo")
		return err
	}

	err = operations.ApplyArgocdDescriptor(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not commit argocd descriptor")
		return err
	}

	err = operations.ApplyArgocdDescriptor(config)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not apply argocd descriptor")
		return err
	}

	time.Sleep(60 * time.Second)

	err = copyRegistrySecret(config.Name)
	if err != nil {
		sendMessage(client, channelID, serviceName, "Could not copy default registry secret")
		log.Printf("Could not copy default registry secret: %v", err)
	}

	sendMessage(client, channelID, serviceName, "Service created!")

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

func copyRegistrySecret(newNamespace string) error {
	err := k8s_api.CopySecret("evilmeow-registry-secret", "owl-of-athena", newNamespace)
	return err
}
