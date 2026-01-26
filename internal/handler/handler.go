package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anil-vinnakoti/newsapi/internal/logger"
	"github.com/google/uuid"
)

type NewsStorer interface {
	// Create news from post request body
	Create(NewsPostRequestBody) (NewsPostRequestBody, error)

	// FindByID news by its ID
	FindByID(uuid.UUID) (NewsPostRequestBody, error)

	// FindAll returns all news in the store
	FindAll() ([]NewsPostRequestBody, error)

	// DeleteByID deletes a news item by its ID
	DeleteByID(uuid.UUID) error

	// UpdateByID updates a news resource by its ID
	UpdateByID(NewsPostRequestBody) error
}

func GetAllNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerFromContext(r.Context())
		logger.Info("request recieved")

		news, err := ns.FindAll()
		if err != nil {
			logger.Error("failed to fertch all news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		allNewsResponse := AllNewsResponse{News: news}
		if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
			logger.Error("failed to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func GetNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerFromContext(r.Context())
		logger.Info("request recieved")

		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("news not found", "newsID", newsID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&news); err != nil {
			logger.Error("failed to encode", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerFromContext(r.Context())
		logger.Info("request recieved")

		var newsRequestBody NewsPostRequestBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(newsRequestBody); err != nil {
			logger.Error("error creating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerFromContext(r.Context())
		logger.Info("request recieved")

		var newsRequestBody NewsPostRequestBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Info("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.UpdateByID(newsRequestBody); err != nil {
			logger.Error("error updating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func DeleteNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerFromContext(r.Context())
		logger.Info("request recieved")

		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id is not a valid uuid", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("news not found", "newsID", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	}
}

type AllNewsResponse struct {
	News []NewsPostRequestBody `json:"news"`
}
