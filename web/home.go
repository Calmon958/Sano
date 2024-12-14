package web

import (
	"net/http"
	"html/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./assets/template/index.html")
			if err != nil {
				http.Error(w, "Error: cannot load login page", http.StatusInternalServerError)
				return
	}

	err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error: cannot load login page", http.StatusInternalServerError)
			return
		}
	
} 