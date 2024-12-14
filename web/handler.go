package web

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type LoginRequest struct {
	PhoneNo string `json:"phone_no"`
	IDNo    string `json:"ID_no"`
}

type OTPResponse struct {
	Message string `json:"message"`
}

func generateOTP() (string, error) {
	max := big.NewInt(1000000)
	otp, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", otp), nil
}

func sendOTP(phoneNumber, otp string) error {
	log.Printf("Sending OTP code %s to phone number %s\n", otp, phoneNumber)
	return nil
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileID := uuid.New().String()
	fileExtension := filepath.Ext(header.Filename)
	fileName := fileID + fileExtension

	driveFileID, err := uploadFileToGoogleDrive(fileName, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = updatePatientRecord(r.FormValue("patientID"), driveFileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully")
}

func uploadFileToGoogleDrive(fileName string, file io.Reader) (string, error) {
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return "", fmt.Errorf("unable to create Drive client: %v", err)
	}

	driveFile := &drive.File{Name: fileName}
	uploadedFile, err := srv.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("unable to upload file to Drive: %v", err)
	}

	return uploadedFile.Id, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var exsist bool
	err := db.QueryRow(`SELECT EXSIST(SELECT 1 FROM users WHERE phone_number=?)`,
		req.PhoneNo, req.IDNo).Scan(&exsist)
	if err != nil {
		http.Error(w, "Error checking user", http.StatusInternalServerError)
		return
	}

	if !exsist {
		http.Error(w, "Invalid Phone number or ID number", http.StatusUnauthorized)
		return
	}

	otp, _ := generateOTP()
	expiry := time.Now().Add(5 * time.Minute)

	_, err = db.Exec(`UPDATE users SET otp_code=?, otp_expiry=? WHERE phone_number=?`, otp, expiry, req.PhoneNo)
	if err != nil {
		http.Error(w, "Error uptdating OTP", http.StatusInternalServerError)
		return
	}

	if err := sendOTP(req.PhoneNo, otp); err != nil {
		http.Error(w, "Error sending OTP", http.StatusInternalServerError)
		return
	}

	response := OTPResponse{
		Message: "OTP sent successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)
}

var upgradeChat = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeChat.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	for {
		var msg struct {
			Content   string
			Encrypted bool
		}

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
			break
		}

		log.Printf("Received: %v", msg.Content)
	}
}

func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating message:", err)
	}
	return privateKey, &privateKey.PublicKey
}

func encryptMessage(publicKey *rsa.PublicKey, message string) []byte {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		log.Fatal("Error encrypting message", err)
	}
	return ciphertext
}

func decryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) string {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		log.Fatal("Error decrypting message:", err)
	}
	return string(plaintext)
}
