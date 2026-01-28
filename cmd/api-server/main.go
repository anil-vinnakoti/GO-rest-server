package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/anil-vinnakoti/newsapi/internal/logger"
	"github.com/anil-vinnakoti/newsapi/internal/news"
	"github.com/anil-vinnakoti/newsapi/internal/postgres"
	"github.com/anil-vinnakoti/newsapi/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	db, err := postgres.NewDB(&postgres.Config{
		Host:     os.Getenv("DATABASE_HOST"),
		DBName:   os.Getenv("DATABASE_NAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		User:     os.Getenv("DATABASE_USER"),
		Port:     os.Getenv("DATABASE_PORT"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Error("db error", "err", err)
		os.Exit(1)
	}
	newsStore := news.NewStore(db)

	r := router.New(newsStore)
	wrappedRoutesHandler := logger.AddLoggerMiddleWare(log, logger.LoggerMiddleware(r))

	log.Info("server starting on port 8080")

	if err := http.ListenAndServe(":8080", wrappedRoutesHandler); err != nil {
		log.Error("failed to start the server:", "error", err)
	}
}
