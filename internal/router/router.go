package router

import (
	"net/http"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
)

func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	// Get all news
	r.HandleFunc("GET /news", handler.GetAllNews(ns))

	// Get news by ID
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID(ns))

	// Create a news
	r.HandleFunc("POST /news", handler.PostNews(ns))

	// Update a news by ID
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID(ns))

	// Delete a new by ID
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID(ns))

	return r
}
