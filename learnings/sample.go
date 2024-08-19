package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Book represents a book with an ID, Title, and Author.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books  = make(map[int]Book)
	nextID = 1
	mu     sync.Mutex
)

// getBooks handles GET requests to retrieve all books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

// getBook handles GET requests to retrieve a book by ID.
func getBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// createBook handles POST requests to create a new book.
func createBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	book.ID = nextID
	nextID++
	books[book.ID] = book
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// updateBook handles PUT requests to update an existing book by ID.
func updateBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	_, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	var updatedBook Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updatedBook.ID = id
	books[id] = updatedBook
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// deleteBook handles DELETE requests to remove a book by ID.
func deleteBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
