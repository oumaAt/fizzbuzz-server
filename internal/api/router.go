package api

import (
	"net/http"

	"github.com/oumaAt/fizzbuzz-server/internal/stats"
)

func NewRouter(store *stats.Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /fizzbuzz", FizzBuzzHandler(store))
	mux.HandleFunc("GET /stats", StatsHandler((store)))
	return mux
}
