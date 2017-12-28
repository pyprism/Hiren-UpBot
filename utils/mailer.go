package utils

import (
	"fmt"
	"log"

	"gopkg.in/mailgun/mailgun-go.v1"
)

func Mail(domain, apiKey, publicAPIKey string) {
	mg := mailgun.NewMailgun(domain, apiKey, publicAPIKey)
	message := mailgun.NewMessage(
		"sender@example.com",
		"Fancy subject!",
		"Hello from Mailgun Go!",
		"recipient@example.com")
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}