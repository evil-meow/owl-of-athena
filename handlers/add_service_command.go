package handlers

import (
	"evil-meow/owl-of-athena/github_api"
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

// handleHelloCommand will take care of /add_service submissions
func HandleAddServiceCommand(command slack.SlashCommand, client *slack.Client) error {
	// The Input is found in the text field so
	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}
	// Add Some default context like user who mentioned the bot
	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: command.UserName,
		},
	}

	serviceName := command.Text

	// Acknowledge that the request was received
	attachment.Text = fmt.Sprintf("Adding the service at %s", serviceName)
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}

	if github_api.IsGithubRepoCreated(serviceName) {
		fmt.Sprintf("Repo %s exists", serviceName)
	} else {
		fmt.Sprintf("Repo %s does not exist", serviceName)
	}

	return nil
}
