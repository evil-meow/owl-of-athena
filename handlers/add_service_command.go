package handlers

import (
	"evil-meow/owl-of-athena/github_api"
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

// handleHelloCommand will take care of /add_service submissions
func HandleAddServiceCommand(command slack.SlashCommand, client *slack.Client) error {

	serviceName := &command.Text
	username := &command.UserName
	channelID := &command.ChannelID

	sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Adding the service at %s", *serviceName))

	if github_api.IsGithubRepoCreated(*serviceName) {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Repo %s exists", *serviceName))
	} else {
		sendMessage(client, username, channelID, serviceName, fmt.Sprintf("Repo %s does not exist. Please, specify an existing repo.", *serviceName))
	}

	// Check if config file exists

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
		fmt.Errorf("failed to post message: %w", err)
	}
}
