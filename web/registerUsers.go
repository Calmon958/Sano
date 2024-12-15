package web

import (
	"github.com/gofrs/uuid"
	"net/http"

)

type GeneralPractioner struct {
	Id             string
	Name           string
	Email          string
	Mobile         string
	IdNumber       string
	Specialization string
	Location       string
	Licence        string
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
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "DoctorRegistration.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var gp GeneralPractioner
		gp.Name = r.FormValue("name")
		gp.Email = r.FormValue("email")
		gp.Mobile = r.FormValue("mobilenumber")
		gp.IdNumber = r.FormValue("idnumber")
		// 	Specialization string //could be more than one
		gp.Licence = r.FormValue("licence")
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Error: generating new uuid", http.StatusInternalServerError)
			return
		}
		gp.Id = id.String()

		//add function to addb


		http.Redirect(w,r,"/LoginGP",http.StatusSeeOther)
	}

}

func RegisterPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "patientRegister.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var patient patient
		patient.Name = r.FormValue("name")
		patient.Mobile = r.FormValue("mobilenumber")
		patient.IdNumber = r.FormValue("idnumber")
		patient.Age = r.FormValue("Age")
		patient.Location = r.FormValue("location")
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Error: generating new uuid", http.StatusInternalServerError)
			return
		}
		patient.Id = id.String()
		//add function to add patient to the database


		//redirect to login page
		http.Redirect(w,r,"/LoginUser",http.StatusSeeOther)
	}



}
