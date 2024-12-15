package web

import (
	"net/http"	
)


func LoginPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "LoginUser.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		//capture login data
		//validate login data
		
		http.Redirect(w,r,"/DashboardPatient",http.StatusSeeOther) 

	}

}

func LoginDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err = templates.ExecuteTemplate(w, "LoginGP.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		//capture login data
		//validate login data
		
		http.Redirect(w,r,"/DashboardDoctor",http.StatusSeeOther) 
	}
}

