package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/anil-vinnakoti/newsapi/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	logger.Info("server starting on port 8080")

	routesHandler := router.New()
	if err := http.ListenAndServe(":8080", routesHandler); err != nil {
		logger.Error("failed to start the server:", "error", err)
	}
}
