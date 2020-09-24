package main

import "net/http"

func addTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}
