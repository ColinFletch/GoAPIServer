package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	port        = ":8080"
	userAPIResp = `
<p> Hello, World! 
<p> Request received : %q endpoint
<p> HTTP Method : %v
<p> This is call %v to this API
`
	countAPIResp = `
<p> Hello, World! 
<p> Request received : %q endpoint
<p> HTTP Method : %v
<p> API Call Count: %v
`
)

var userCount int

type countHandler struct {
	mu sync.Mutex //guards nil
	n  int
}

func userHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	userCount++
	out := fmt.Sprintf(userAPIResp, r.URL.Path, r.Method, userCount)
	fmt.Fprintf(w, out)
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	out := fmt.Sprintf(countAPIResp, r.URL.Path, r.Method, h.n)
	fmt.Fprintf(w, out)
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.HandleFunc("/users/", userHandlerFunc)
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(port, nil))
}
