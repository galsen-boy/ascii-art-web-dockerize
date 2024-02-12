package main

import (
	"log"
	"net/http"
	"text/template"
)

// Handles 400 bad request
func error400(w http.ResponseWriter) {
	tpl400, err := template.ParseFiles("./templates/400.page.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(400)
	tpl400.Execute(w, nil)
}

// Handles 404 not found
func error404(w http.ResponseWriter) {
	tpl404, err := template.ParseFiles("./templates/404.page.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(404)
	tpl404.Execute(w, nil)
}

// Handles 500 internal server error
func error500(w http.ResponseWriter) {
	tpl500, err := template.ParseFiles("./templates/500.page.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(500)
	tpl500.Execute(w, nil)
}

// method not allowed
func error405(w http.ResponseWriter) {
	tpl405, err := template.ParseFiles("./templates/405.page.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(405)
	tpl405.Execute(w, nil)
}
