package main

import (
	"fmt"
	"log"
)

type Email struct {
	ID      int
	Address string
	Subject string
	Body    string
	Sent    int
}

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to the database")
	}
	defer db.Close()

	rows, err := db.Query(`select * from outbox `)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var allMsg []Email

	for rows.Next() {
		var msg Email
		if err := rows.Scan(&msg.ID, &msg.Address, &msg.Subject, &msg.Body, &msg.Sent); err != nil {
			log.Print(err)
		}
		allMsg = append(allMsg, msg)
		// NOTE: sanity check
		fmt.Printf(
			"sent?: %d\taddr : %s\nsub: %s\t\nbody: %s\n---------\n",
			msg.Sent,
			msg.Address,
			msg.Subject,
			msg.Body,
		)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}
}
