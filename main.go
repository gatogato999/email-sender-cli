package main

import (
	"log"

	"github.com/joho/godotenv"
)

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

	ConcurrentSend(db, unSentMsgs)
}
