package web

import (
	"net/http"
)

func PatientPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "PatientDashboard.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
}

func PatientRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "Records.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
}

func PatientChat(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "chat.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
}

func PatientDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "MedicalDocuments.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
}

func PatientProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "PatientProfile.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
}


