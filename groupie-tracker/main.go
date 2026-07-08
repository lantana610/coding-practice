package main

import (
	"fmt"
	"net/http"
)

var allArtists []FullArtist

func main() {
	artists, err := getAllArtists()
	if err != nil {
		fmt.Println("error fetching data:", err)
		return
	}
	allArtists = artists

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/artist/", artistsHandler)
	http.ListenAndServe(":8080", nil)
}
