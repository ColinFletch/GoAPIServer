package handler

import (
	"encoding/json"
	"github.com/colinfletch/GoAPIServer/platform/newsfeed"
	"net/http"
)

// NewsFeedGet Comment
func NewsFeedGet(feed newsfeed.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
