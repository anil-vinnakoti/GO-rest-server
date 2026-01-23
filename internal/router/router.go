package router

import (
	"net/http"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Get all news
	r.HandleFunc("GET /news", handler.GetAllNews())

	// Get news by ID
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID())

	// Create a news
	r.HandleFunc("POST /news", handler.PostNews())

	// Update a news by ID
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID())

	// Delete a new by ID
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID())

	return r
}
