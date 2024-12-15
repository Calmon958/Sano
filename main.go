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

	http.HandleFunc("/loginPatient", handler.LoginPatient)
	http.HandleFunc("/loginDoctor", handler.LoginDoctor)
	http.HandleFunc("/", handler.HomeHandler)
	
	http.HandleFunc("/DocReg", handler.RegisterGp)
	http.HandleFunc("/PatientReg", handler.RegisterPatient)
	
	http.HandleFunc("/DashboardPatient", handler.PatientPageHandler)
	http.HandleFunc("/DashboardDoctor", handler.DoctorPageHandler)
	http.HandleFunc("/DoctorRecord", handler.DoctorRecords)
	http.HandleFunc("/DoctorDocuments", handler.DoctorMedicalDocs)
	
	
	http.HandleFunc("/PatientRecord", handler.PatientRecord)
	http.HandleFunc("/PatientChat", handler.PatientChat)
	http.HandleFunc("/PatientDocuments", handler.PatientDocument)
	http.HandleFunc("/PatientProfile", handler.PatientProfile)
	fmt.Println("Server is running on http://localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
