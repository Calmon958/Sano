package web

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

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
