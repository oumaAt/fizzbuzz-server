package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/oumaAt/fizzbuzz-server/internal/fizzbuzz"
	"github.com/oumaAt/fizzbuzz-server/internal/stats"
)

const MaxLimit = 10000

type params struct {
	int1, int2, limit int
	str1, str2        string
}

func FizzBuzzHandler(store *stats.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := parseParams(r.URL.Query())
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := fizzbuzz.Generate(p.int1, p.int2, p.limit, p.str1, p.str2)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		store.Record(stats.Request{
			Int1:  p.int1,
			Int2:  p.int2,
			Limit: p.limit,
			Str1:  p.str1,
			Str2:  p.str2,
		})
		writeJSON(w, http.StatusOK, result)
	}
}

func parseParams(q url.Values) (params, error) {
	var p params
	var err error

	if p.int1, err = parseIntParam(q, "int1"); err != nil {
		return params{}, err
	}

	if p.int2, err = parseIntParam(q, "int2"); err != nil {
		return params{}, err
	}

	if p.limit, err = parseIntParam(q, "limit"); err != nil {
		return params{}, err
	}

	if p.limit > MaxLimit {
		return params{}, fmt.Errorf("limit must not exceed %d", MaxLimit)
	}

	p.str1, p.str2 = q.Get("str1"), q.Get("str2")
	if p.str1 == "" || p.str2 == "" {
		return params{}, fmt.Errorf("str1 and str2 are required")
	}

	return p, nil
}

func parseIntParam(q url.Values, name string) (int, error) {
	raw := q.Get(name)
	if raw == "" {
		return 0, fmt.Errorf("%s is required", name)
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", name)
	}
	return n, nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
