package handler

import (
	"encoding/json"
	"github.com/colinfletch/GoAPIServer/platform/newsfeed"
	"net/http"
)

//NewsfeedPost Comment
/* Post to newsfeed in a browser console simply by using:
await fetch('/news', {
	method: 'POST',
	headers: { 'content-type': 'application/json' },
	body: JSON.stringify({
		title: 'another',
		post: 'one'
			})
})
*/
func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})
		w.WriteHeader(201)
		w.Write([]byte("Well Done!"))
	}
}
