package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Text   string
	Result string
	Banner string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var data PageData

	// tmpl, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
