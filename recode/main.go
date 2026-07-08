package main

import "net/http"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", handleAsciiArt)
	http.HandleFunc("/ascii-art-switch", handleSwitch)

	http.ListenAndServe(":8080", nil)
}
