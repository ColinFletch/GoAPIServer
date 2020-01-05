package handler

import (
	"encoding/json"
	"github.com/colinfletch/GoAPIServer/platform/newsfeed"
	"net/http"
)

// NewsfeedGet Comment
func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
