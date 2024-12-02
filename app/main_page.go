package main

import (
	"net/http"
	"text/template"
)

// mainPage serves the main HTML page
func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Not Found - Template file missing", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK) // Explicitly set status code 200
	tmpl.Execute(w, nil)
}
