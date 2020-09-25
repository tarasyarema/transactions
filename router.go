package main

import "net/http"

func (app *App) transactionRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Testing purposes
	case http.MethodGet:
		app.getTransactions(w, r)
		return
	case http.MethodPost:
		app.addTransaction(w, r)
		return
	case http.MethodDelete:
		app.deleteTransactions(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (app *App) statisticsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getStatistics(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
