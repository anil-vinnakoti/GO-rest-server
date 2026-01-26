package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/anil-vinnakoti/newsapi/internal/logger"
	"github.com/anil-vinnakoti/newsapi/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	routesHandler := router.New(nil)
	wrappedRoutesHandler := logger.AddLoggerMiddleWare(log, logger.LoggerMiddleware(routesHandler))

	log.Info("server starting on port 8080")

	if err := http.ListenAndServe(":8080", wrappedRoutesHandler); err != nil {
		log.Error("failed to start the server:", "error", err)
	}
}
