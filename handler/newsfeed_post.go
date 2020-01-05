package handler

import (
	"encoding/json"
	"net/http"

	"github.com/colinfletch/GoAPIServer/platform/newsfeed"
)

//NewsFeedPost Comment
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
func NewsFeedPost(feed *newsfeed.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)
		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})
		w.Write([]byte("Well Done!"))
	}
}
