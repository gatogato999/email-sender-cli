package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func SendEmail(recept, subj, body string, cntx context.Context) error {
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

func ConccurtSend(db *sql.DB, unSentMsgs []Email) {
	var waitter sync.WaitGroup
	results := make(chan string, len(unSentMsgs))

	for _, value := range unSentMsgs {
		waitter.Add(1)

		go func(msg Email) {
			defer waitter.Done()
			cntx, cancelAfterAMin := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cancelAfterAMin()

			if sendError := SendEmail(
				value.Address,
				value.Subject,
				value.Body,
				cntx,
			); sendError != nil {
				results <- fmt.Sprintf("can't send msg %d : %v", value.ID, sendError)
			} else {
				if changeStateError := MarkAsSent(db, value.ID); changeStateError != nil {
					results <- fmt.Sprintf("msg %d sent but not set to read : %v\n", value.ID, changeStateError)
				} else {
					results <- fmt.Sprintf("msg %d sent and set to read \n", value.ID)
				}
			}
		}(value)
	}

	// Wait and close results
	go func() {
		waitter.Wait()
		close(results)
	}()

	// 5. Consume messages
	for msg := range results {
		fmt.Println(msg)
	}
}
