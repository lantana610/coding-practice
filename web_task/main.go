package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}
type User struct {
	ID    int
	Name  string
	Email string
}
type Use struct {
	Name string
	Age  int
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	term, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "don't panic error coming from the server not you", http.StatusInternalServerError)
		return
	}
	books := []Book{
		{Title: "Go Programming", Author: "John Doe"},
		{Title: "Learning HTTP", Author: "Jane Smith"},
		{Title: "Mastering APIs", Author: "Alex Brown"},
	}
	term.Execute(w, books)
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/greet" {
		http.Error(w, "wrong path page not found", http.StatusNotFound)
		return
	}

	input := r.URL.Query().Get("name")
	if input == "" {
		input = "Guest"
	}
	fmt.Fprintf(w, "hello, %v!", input)
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.Error(w, "worng path cound not find this page", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	input := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	json.NewEncoder(w).Encode(input)
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var user Use
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User %s registered successfully. Age: %d", user.Name, user.Age)
}

func main() {
	http.HandleFunc("/books", bookHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("server runing at http://localhost:8080/books")
	fmt.Println("http://localhost:8080/greet?name=Alice")
	fmt.Println("http://localhost:8080/greet")
	fmt.Println("http://localhost:8080/user")
	fmt.Println("server at http://localhost:8080/register")

	http.ListenAndServe(":8080", nil)
}
