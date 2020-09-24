package main

import "net/http"

func getStatistics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}
