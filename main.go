package main

import (
	"fmt"
	"log"
	"net/http"

	db "Sano/database"
	handler "Sano/web"
)

func main() {
	db.InitDB()
	db.CreateTable()

	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/", handler.HomeHandler)
	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
