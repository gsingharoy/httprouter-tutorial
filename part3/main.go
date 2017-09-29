package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/books", BookIndex)
	router.GET("/books/:isdn", BookShow)

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
