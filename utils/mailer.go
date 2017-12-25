package utils

import  "gopkg.in/mailgun/mailgun-go.v1"

mg := mailgun.NewMailgun(yourdomain, ApiKey, publicApiKey)
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