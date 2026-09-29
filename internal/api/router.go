package api

import (
	"net/http"

	"github.com/oumaAt/fizzbuzz-server/internal/stats"
)

func NewRouter(store *stats.Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /fizzbuzz", FizzBuzzHandler(store))
	mux.HandleFunc("GET /stats", StatsHandler((store)))
	mux.HandleFunc("GET /health", healthHandler)
	return mux
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
