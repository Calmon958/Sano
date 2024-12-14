package web

import (
	"github.com/gofrs/uuid"
	"net/http"
	"fmt"
	"Sano/database"
	_ "github.com/mattn/go-sqlite3"
)

type GeneralPractioner struct {
	Id             string
	firstName           string
	lastName           string
	year_of_birth          string
	email          string
	phone_number        string
	doctor_id_number       string
	title string
	location       string
	license_number       string
}

type patient struct {
	Id       string
	Name     string
	IdNumber string
	Mobile   string
	Age      string
	Location string
}

func RegisterGp(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodGet{
		err = templates.ExecuteTemplate(w, "doctorRegistration.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var gp GeneralPractioner
		gp.firstName = r.FormValue("first_name")
		gp.lastName = r.FormValue("last_name")
		gp.email = r.FormValue("email")
		gp.year_of_birth = r.FormValue("year_of_birth")
		gp.doctor_id_number = r.FormValue("doctor_id_number")
		gp.license_number = r.FormValue("license_number")
		gp.location = r.FormValue("location")
		gp.phone_number = r.FormValue("phone_number")
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Error: generating new uuid", http.StatusInternalServerError)
			return
		}
		gp.Id = id.String()

		//fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s\n", gp.firstName, gp.lastName, gp.email, gp.phone_number, gp.year_of_birth, gp.doctor_id_number,gp.location, gp.license_number)
		query := `
		INSERT INTO  doctor_registration(first_name, last_name, email, year_of_birth, doctor_id_number,license_number,location, phone_number) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`
		_, err = db.DB.Exec(query, gp.firstName, gp.lastName, gp.email, gp.year_of_birth, gp.doctor_id_number, gp.license_number,gp.location, gp.phone_number)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Registration successful! Welcome, %s", gp.firstName)
	}
}


func RegisterPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var patient patient
		patient.Name = r.FormValue("firstName")
		patient.Mobile = r.FormValue("phone")
		patient.IdNumber = r.FormValue("idNumber")
		patient.Age = r.FormValue("yob")
		patient.Location = r.FormValue("location")
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Error: generating new uuid", http.StatusInternalServerError)
			return
		}
		patient.Id = id.String()
	}

	//add function to add patient to the database
}
