package main

type Book struct {
	// The main identifier for the Book. This will be unique.
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// A map to store the books with the ISDN as the key
// This acts as the storage in lieu of an actual database
var bookstore = make(map[string]*Book)
