package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Printf("Starting server...\n")
	books = append(books, Book{ID: "1", Title: "Lord of Rings"})

	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
