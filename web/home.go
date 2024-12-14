package web

import (
	"net/http"
	"html/template"
	"log"
)

var templates *template.Template
var err error

func init() {
	templates, err = template.ParseGlob("./assets/template/*.html")
	if err != nil {
		log.Fatalf("Failed to render template")
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		http.Redirect(w,r,"doctor_registration", 302)
	}
} 