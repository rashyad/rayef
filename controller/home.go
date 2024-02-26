package controller

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("view/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
