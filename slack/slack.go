package slack

import (
	"fmt"
	"go-mailapp/config"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

// MakeNotification ...
func MakeNotification(message string) bool {
	payload := slack.Payload{
		Text: message,
	}

	err := slack.Send(config.SlackWebhookURL, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}

	return true

}
