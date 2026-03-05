package main

import (
	"log"

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

	unSentMsgs, msgFetchError := GetAllMsg(db)
	if msgFetchError != nil {
		log.Fatal(msgFetchError)
	}

	for _, value := range unSentMsgs {
		if sendError := SendEmail(value.Address, value.Subject, value.Body); sendError != nil {
			log.Printf("can't send msg %d : %v", value.ID, sendError)
		}
	}
}
