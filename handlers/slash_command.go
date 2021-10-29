package handlers

import (
	"log"

	"github.com/slack-go/slack"
)

// handleSlashCommand will take a slash command and route to the appropriate function
func HandleSlashCommand(command slack.SlashCommand, client *slack.Client) error {
	// We need to switch depending on the command
	switch command.Command {
	case "/add_service":
		// This was a hello command, so pass it along to the proper function
		err := HandleAddServiceCommand(command, client)
		if err != nil {
			log.Printf("Error handling command: %s", err)
		}
	}

	return nil
}
