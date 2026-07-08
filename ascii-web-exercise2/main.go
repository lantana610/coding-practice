package main

import (
	"fmt"
	"io"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "You made a GET request.", http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		http.Error(w, "You made a POST request.", http.StatusOK)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "You made a [METHOD] request.", http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "You made a [METHOD] request.", http.StatusMethodNotAllowed)
		return
	}
}
func echohandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed only post method allowe", http.StatusMethodNotAllowed)
		return
	}
	text, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}
	r.Body.Close()
	if len(text) == 0 {
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
		return
	}
	w.Write(text)

}
func header(w http.ResponseWriter, r *http.Request) {
	text := Header.Get("X-Custom-Token")
	if len(text) == 0 {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	if len(text) != 0 {
		fmt.Fprint(w, "Token received: abc123")
	}

}

func main() {
	http.HandleFunc("/method-inspector", methodHandler)
	http.HandleFunc("/echo", echohandler)
	http.HandleFunc("/headers", header)

	fmt.Println("server runing http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
