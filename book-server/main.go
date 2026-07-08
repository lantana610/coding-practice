package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	Id             string `json:"id" `
	Title          string `json:"title" `
	Author         string `json:"author" `
	Published_year int    `json:"published_year" `
	Publisher      string `json:"publisher" `
	Genre          string `json:"genre" `
	ISBN           string `json:"isbn" `
	Language       string `json:"langauge"`
}

var bookstore = map[string]Book{
	"Book1": {Id: "uuid-1",
		Title:          "Go programming made simple",
		Author:         "Lantana Yusuf",
		Published_year: 2025,
		Publisher:      "Codec Publishing",
		Genre:          "Tech",
		ISBN:           "0987654334567",
		Language:       "English"},

	"Book2": {Id: "uuid-2",
		Title:          "Go programming made simple",
		Author:         "Yaks Nicodemus",
		Published_year: 2025,
		Publisher:      "Codec Publishing",
		Genre:          "Tech",
		ISBN:           "0987654334567",
		Language:       "English"},
}

// func homeHandler(w http.ResponseWriter, req *http.Request) {

// }
func bookHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	books := make([]Book, 0, len(bookstore))
	for _, b := range bookstore {
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
}
func GetBookHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "books" {
		http.NotFound(w, req)
	}
	id := parts[2]

	book, found := bookstore[id]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "book not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
func createBookHandler(w http.ResponseWriter, req *http.Request){
	
}
func main() {
	http.HandleFunc("/books", bookHandler)
	http.HandleFunc("/books/", GetBookHandler)
	fmt.Println("server runing at http://localhost:8080/books/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error")
	}

}
