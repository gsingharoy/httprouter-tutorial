package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestBookShow(t *testing.T) {
	t.Log("When the books' isdn does not exist")
	// A request with a non-existant isdn
	req1, err := http.NewRequest("GET", "/books/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr1 := newRequestRecorder(req1, "GET", "/books/:isdn", BookShow)
	if rr1.Code != 404 {
		t.Error("Expected response code to be 404")
	}
	// expected response
	er1 := "{\"error\":{\"status\":404,\"title\":\"Record Not Found\"}}\n"
	if rr1.Body.String() != er1 {
		t.Error("Response body does not match")
	}

	t.Log("When the book exists")
	// Create an entry of the book to the bookstore map
	testBook := &Book{
		ISDN:   "111",
		Title:  "test title",
		Author: "test author",
		Pages:  42,
	}
	bookstore["111"] = testBook
	// A request with an existing isdn
	req2, err := http.NewRequest("GET", "/books/111", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := newRequestRecorder(req2, "GET", "/books/:isdn", BookShow)
	if rr2.Code != 200 {
		t.Error("Expected response code to be 200")
	}
	// expected response
	er2 := "{\"meta\":null,\"data\":{\"isdn\":\"111\",\"title\":\"test title\",\"author\":\"test author\",\"pages\":42}}\n"
	if rr2.Body.String() != er2 {
		t.Error("Response body does not match")
	}
}

func TestBookIndex(t *testing.T) {
	// Create an entry of the book to the bookstore map
	testBook := &Book{
		ISDN:   "111",
		Title:  "test title",
		Author: "test author",
		Pages:  42,
	}
	bookstore["111"] = testBook
	// A request with an existing isdn
	req1, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "GET", "/books", BookIndex)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}
	// expected response
	er1 := "{\"meta\":null,\"data\":[{\"isdn\":\"111\",\"title\":\"test title\",\"author\":\"test author\",\"pages\":42}]}\n"
	if rr1.Body.String() != er1 {
		t.Error("Response body does not match")
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}
