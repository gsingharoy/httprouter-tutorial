package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

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

// Handler for the books Show action
// GET /books/:isdn
func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	book, ok := bookstore[isdn]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !ok {
		// No book with the isdn in the url has been found
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{Error: &apiError{status: 404, title: "Record Not Found"}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}
	response := JsonResponse{Data: book}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
