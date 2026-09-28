package api

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /fizzbuzz", FizzBuzzHandler)
	return mux
}