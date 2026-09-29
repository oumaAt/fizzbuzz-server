package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/oumaAt/fizzbuzz-server/internal/api"
	"github.com/oumaAt/fizzbuzz-server/internal/stats"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := stats.NewStore()

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           api.NewRouter(store),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
