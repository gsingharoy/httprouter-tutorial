package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

type Book struct {
	// The main identifier for the Book. This will be unique.
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// A map to store the books with the ISDN as the key
// This acts as the storage in lieu of an actual database
var bookstore = make(map[string]*Book)

// Handler for the books index action
// GET /books
func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range bookstore {
		books = append(books, book)
	}
	response := &JsonResponse{Data: &books}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/books", BookIndex)

	// Create a couple of sample Book entries
	bookstore["123"] = &Book{
		ISDN:   "123",
		Title:  "Silence of the Lambs",
		Author: "Thomas Harris",
		Pages:  367,
	}

	bookstore["124"] = &Book{
		ISDN:   "124",
		Title:  "To Kill a Mocking Bird",
		Author: "Harper Lee",
		Pages:  320,
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
