package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

)

var db *sql.DB

func init(){
	var err error

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil{
		log.fatal("Error connecting to database: %v\n", err)
	}


	createTable := "
	CREATE TABLE IF NOT EXSISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		phone_number TEXT NOT NULL UNIQUE,
		ID_Number TEXT NOT NULL,
		otp_code TEXT,
		otp_expiry DATETIME,
	)"

	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}
}

type LoginRequest struct {
	PhoneNo string `json:"phone_no"`
	IDNo string `json:"ID_no"`
}

type OTPResponse struct {
	Message string `json:"message"`
}

//generate otp code
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

//send OTP code generated
func sendOTP(phoneNumber, otp string) error {
	log.Printf("Sending OTP code %s to phone number %s\n", otp, phoneNumber)
	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http,StatusBadRequest)
		return
	}

	var exsist bool
	err := db.QueryRow(`SELECT EXSISTS(SELECT 1 FROM users WHERE phone_number=?)`,
	otp, expiry, req.PhoneNo).Scan(&exsist)
	if err != nil {
		http.Error(w, "Error checking user", http,StatusInternalServerError)
		return
	}

	if !exsist {
		http.Error(w, "Invalid Phone number or ID number", http,StatusUnauthorized)
		return
	}

	otp := generateOTP()
	expiry := time.Now().Add(5 *time.Minutes)

	_, err := db.Exec(`UPDATE users SET otp_code=?, otp_expiry=? WHERE phone_number=?`,
	otp, expiry, req.PhoneNo)
	if err != nil {
		http.Error(w, "Error uptdating OTP", http,StatusInternalServerError)
		return
	}


	if err := sendOTP(req.PhoneNo, otp); err != nil {
		http.Error(w, "Error sending OTP", http,StatusInternalServerError)
		return
	}

	response :=OTPResponse{
		Message: "OTP sent successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(w).Encode(responsse)

	http.Redirect(w, r, "/Dashboard", http.StatusSeeOthers)
}

func main(){
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServer(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}