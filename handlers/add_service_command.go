package handlers

import (
	"evil-meow/owl-of-athena/github_api"
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

	sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Adding the service at %s", *serviceName))

	if github_api.IsGithubRepoCreated(*serviceName) {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Repo %s exists", *serviceName))
	} else {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Repo %s does not exist. Please, specify an existing repo.", *serviceName))
	}

	_, err := readConfigFile(serviceName)
	if err != nil {
		sendMessage(client, username, channelID, serviceName, "Could not find owl.yml at the root of the repo. Please, create it in order to add the service.")
		return err
	}

	infraRepoName := *serviceName + "-infra"

	if github_api.IsGithubRepoCreated(infraRepoName) {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Infra repo %s already exists", infraRepoName))
	} else {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Infra repo %s does not exist. Creating.", infraRepoName))
		github_api.CreateGitubRepo(&infraRepoName)
	}

	return nil
}

func sendMessage(client *slack.Client, username *string, channelID *string, serviceName *string, text string) {
	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: *username,
		}, {
			Title: "Service",
			Value: *serviceName,
		},
	}

	attachment.Text = text
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(*channelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Printf("failed to post message: %w", err)
	}
}

func readConfigFile(serviceName *string) (string, error) {
	configFileUrl := fmt.Sprintf("https://raw.github.com/evil-meow/%s/main/owl.yml", *serviceName)
	configFile, err := github_api.ReadFile(&configFileUrl)
	if err != nil {
		log.Printf("Could not find config file at: %s", configFileUrl)
		return "", err
	}

	return string(configFile), nil
}
