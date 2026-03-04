package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type Email struct {
	ID      int
	Address string
	Subject string
	Body    string
	Sent    int
}

func main() {
	log.SetFlags(log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to the database")
	}
	defer db.Close()

	emailServer := os.Getenv("SMPTSERVER")
	emailServerPort := os.Getenv("EMAILPORT")
	senderEmail := os.Getenv("EMAILUSERNAME")
	senderPassword := os.Getenv("EMAILPASSWORD")

	auth := smtp.PlainAuth("", senderEmail, senderPassword, emailServer)

	err = smtp.SendMail(
		emailServer+":"+emailServerPort,
		auth,
		senderEmail,
		[]string{"mhmdmrhsn13@gmail.com"},
		[]byte("hi"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
