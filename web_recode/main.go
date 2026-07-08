package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "wrong path page not found", http.StatusNotFound)
		return
	}
	tem, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "not your fault internal ser error", http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	term, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "not your fault internal ser error", http.StatusInternalServerError)
		return
	}
	text := r.FormValue("text")
	term.Execute(w, text)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/echo", echoHandler)

	fmt.Println("server runing at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
