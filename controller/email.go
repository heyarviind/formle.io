package controller

import (
	"fmt"
	"formle/config"

	"github.com/keighl/postmark"
)

//SendEmail ...
func SendEmail(toReplyEmail, toEmail, subject, body string) bool {
	client := postmark.NewClient(config.ServerToken, config.AccountToken)
	email := postmark.Email{
		From:       "no-reply@uicard.io",
		To:         toEmail,
		ReplyTo:    toReplyEmail,
		Subject:    subject,
		HtmlBody:   body,
		TextBody:   body,
		Tag:        "pw-reset",
		TrackOpens: true,
	}

	_, err := client.SendEmail(email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
