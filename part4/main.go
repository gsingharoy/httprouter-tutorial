package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
