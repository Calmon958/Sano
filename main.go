package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handler "Sano/web"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		phone_number TEXT NOT NULL UNIQUE,
		ID_Number TEXT NOT NULL,
		otp_code TEXT,
		otp_expiry DATETIME
	);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/login", handler.LoginHandler)

	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
