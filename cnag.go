package main

import (
	"cnag/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("http server implementation.\n")

	// respond to index path (GET requests)
	http.HandleFunc("/", index)

	// http://localhost:8080/api/books
	http.HandleFunc("/api/books", api.BooksHandleFunc)

	// http://localhost:8080/api/books/1491941197
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	// start the server at port().. and listen
	http.ListenAndServe(port(), nil)

}

// port() returns default port 8080 if it's empty
// can be changed in command line: export PORT=XXXX
func port() string {
	// get the current environment variable PORT
	port := os.Getenv("PORT")

	// default it to 8080 if a port is not specified in the GET request
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

// index function triggered if the url points to /
func index(w http.ResponseWriter, r *http.Request) {
	// StatusOK = http 200
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index Page.")
}
