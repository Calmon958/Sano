package main

import (
	"database/sql"
	"encoding/json"
	"Sano/database"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

)

var db *sql.DB

func init() {
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}
}

func main(){
	db.InitDB()
	db.CreateTable()
	http.HandleFunc("/login", handler.LoginHandler)

	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}