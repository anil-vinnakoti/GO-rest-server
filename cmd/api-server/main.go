package main

import (
	"log"
	"net/http"

	"github.com/anil-vinnakoti/newsapi/internal/router"
)

func main() {
	routesHandler := router.New()
	if err := http.ListenAndServe(":8080", routesHandler); err != nil {
		log.Fatal("failed to start the server:", err)
	}
}
