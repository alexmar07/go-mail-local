package command

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type Sender struct {
	options Options
}

func (s *Sender) Run(args []string) error {

	setup, err := getSetup(s.options["profile"].(string))

	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", setup["from"].(string))

	data := getMailDataFromUser()

	// // Set E-Mail receivers
	m.SetHeader("To", data["to"].(string))

	// // Set E-Mail subject
	m.SetHeader("Subject", data["subject"].(string))

	// // Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", data["body"].(string))

	// Settings for SMTP server
	d := gomail.NewDialer(setup["host"].(string), int(setup["port"].(float64)), setup["user"].(string), setup["password"].(string))

	fmt.Println("Sending email...")

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Invio con successo a " + data["to"].(string))

	return nil
}

func getMailDataFromUser() map[string]interface{} {

	var to string
	var subject string
	var body string

	fmt.Println("Inserisci l'indirizzo email del destinatario:")
	fmt.Scanln(&to)
	fmt.Println("Inserisci il soggetto della mail:")
	fmt.Scanln(&subject)
	fmt.Println("Inserisci il corpo della mail:")
	fmt.Scanln(&body)

	return map[string]interface{}{
		"to":      to,
		"subject": subject,
		"body":    body,
	}
}
