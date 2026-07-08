package main

import (
	"net/http"
)

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	var data PageData
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// if r.Method == http.MethodPost {
	// 	data.IsBool = true
	// }
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	data.Text = text
	data.Banner = banner
	font, err := BannerLoader(banner)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	data.Result = BuildArt(text, font)

	// tmpl, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
