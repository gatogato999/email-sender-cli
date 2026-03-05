package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dbPassword := os.Getenv("DBPASS")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")

	databaseSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", databaseSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllMsg(db *sql.DB) ([]Email, error) {
	rows, queryErr := db.Query(`select * from outbox where sent = 0 `)
	if queryErr != nil {
		log.Println(queryErr)
		return []Email{}, queryErr
	}

	defer rows.Close()

	var allMsg []Email

	for rows.Next() {
		var msg Email
		if scanErr := rows.Scan(
			&msg.ID,
			&msg.Address,
			&msg.Subject,
			&msg.Body,
			&msg.Sent,
		); scanErr != nil {
			if scanErr == rows.Err() {
				log.Printf("error on msg %d : %v", msg.ID, scanErr)
			} else {
				return []Email{}, scanErr
			}
		}
		allMsg = append(allMsg, msg)
	}
	return allMsg, nil
}

func MarkAsSent(db *sql.DB, id int) error {
	_, err := db.Exec(`UPDATE outbox SET sent = 1 WHERE id = ?`, id)
	return err
}
