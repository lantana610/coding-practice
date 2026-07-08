package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
	if r.URL.Path != "/ping" {
		http.Error(w, "wrong path", http.StatusMethodNotAllowed)
		return
	}
}
func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("name")
	if text == "" {
		text = "Guest"
	}
	fmt.Fprintf(w, "Hello, %v!", text)

}
func count(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "Send a POST request with text to count words", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodPost {
		char, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading text", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, len(char))
	}
}
func calculate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculate" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	operation := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	first, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	second, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	// operation == (add) || (subtract) || (multiply)
	switch operation {
	case "add":
		fmt.Fprintf(w, "output %d", first+second)
	case "subtract":
		fmt.Fprintf(w, "output %d", first-second)
	case "multiply":
		fmt.Fprintf(w, "output %d", first*second)
	default:
		http.Error(w, "operation not found", http.StatusBadRequest)
	}

}
func agent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Header)

}
func dash(w http.ResponseWriter, r *http.Request){
	text := r.Header.Get("X-API-Key")
	tex := "secret123"
if tex != text {
	http.Error(w, "unathorised", http.StatusUnauthorized)
	return
}	
fmt.Fprint(w, "Welcome")

}
func legacy(w http.ResponseWriter, r *http.Request){
http.Redirect(w,r, "/v2", http.StatusMovedPermanently)
}
func v2(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/count", count)
	http.HandleFunc("/calculate", calculate)
	http.HandleFunc("/agent", agent)
	http.HandleFunc("/dashboard", dash)
	http.HandleFunc("/legacy", legacy)
	http.HandleFunc("/v2", v2)
	fmt.Println("server runing at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
