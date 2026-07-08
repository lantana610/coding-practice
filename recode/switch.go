package main

import (
	"net/http"
)

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	var data PageData
	// data.IsBool = true
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

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
