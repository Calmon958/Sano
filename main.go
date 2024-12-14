package main

import (
	"fmt"
	"log"
	"net/http"

	handler "Sano/web"

)

func main() {
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/", handler.HomeHandler)
	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
