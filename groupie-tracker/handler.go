package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "error parsing file", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, allArtists)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Path

	parts := strings.Split(data, "/")

	conv, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid artist id", http.StatusBadRequest)

		return
	}
	found := false
	for _, fa := range allArtists {
		if fa.Artist.ID == conv {
			found = true
			tmpl, err := template.ParseFiles("template/artist.html")
			if err != nil {
				http.Error(w, "error parsing artist file", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, fa)
		}

	}
	if !found {
		http.Error(w, "artist not found", http.StatusNotFound)
		return
	}

}
