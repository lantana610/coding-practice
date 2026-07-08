package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandeler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not FOund", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500", http.StatusInternalServerError)
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result := "The text is " + text + " and the banner is: " + banner
	tem.Execute(w, result)
}

func main() {
	http.HandleFunc("/", homeHandeler)
	http.HandleFunc("/ascii-art", asciiHandler)
	fmt.Println("server is runing at  http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
