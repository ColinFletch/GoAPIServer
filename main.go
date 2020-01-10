package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/colinfletch/goapiserver/platform/newsfeed"
	"github.com/go-chi/chi"
)

const (
	port        = ":3000"
	userAPIResp = `
 Request received : %q endpoint
 HTTP Method : %v
 This is call %v to this API
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
	http.HandleFunc("/users/", userHandlerFunc) //allows /users/deadpool etc
	http.Handle("/count", new(countHandler))    // strict matching, only on exact endpoint

	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Something New!",
		Post:  "This is a new article post...",
	})

	r := chi.NewRouter()

	//Get news
	r.Get("/news", func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	})

	//Post news
	r.Post("/news", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})
		w.WriteHeader(201)
		w.Write([]byte("Well Done!"))
	})

	fmt.Println("Serving on port ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
