package api

import (
	"net/http"

	"github.com/oumaAt/fizzbuzz-server/internal/stats"
)

func StatsHandler(store *stats.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		topReq, hits, ok := store.MostFrequent()
		if !ok {
			writeJSON(w, http.StatusOK, map[string]string{"message": "no requests recorded yet"})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"request": topReq,
			"hits":    hits,
		})
	}
}