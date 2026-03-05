package main

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(recept, subj, body string) error {
	emailServer := os.Getenv("SMPTSERVER")
	emailServerPort := os.Getenv("EMAILPORT")
	from := os.Getenv("EMAILUSERNAME")
	pass := os.Getenv("EMAILPASSWORD")

	msg := "From: " + from + "\n" +
		"To: " + recept + "\n" +
		"Subject:" + subj + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, pass, emailServer)
	err := smtp.SendMail(emailServer+":"+emailServerPort,
		auth,
		from, []string{recept}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return nil
}
