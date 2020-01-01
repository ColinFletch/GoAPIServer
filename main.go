package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/users", userHandlerFunc)
	log.Fatal(http.ListenAndServe(port, nil))
}

func userHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Request received : /users endpoint")
	fmt.Fprintf(w, "Response logged for API HTTP Method '%v'", r.Method)
}
