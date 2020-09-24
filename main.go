package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

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

func (app *App) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%v %v %v ms", r.Method, r.RequestURI, time.Since(now).Milliseconds())
	})
}

func main() {
	// The http port that the API will listen to
	port := 8080
	mux := http.NewServeMux()

	app := &App{
		Transactions: make([]Transaction, 0),
	}

	// Define the http handlers
	mux.HandleFunc("/transactions", app.transactionRouter)
	mux.HandleFunc("/statistics", app.statisticsRouter)

	// Apply middleware to the requests
	afterMux := app.middleware(mux)

	portStr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(portStr, afterMux); err != nil {
		log.Fatal(err)
	}
}
