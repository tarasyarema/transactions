package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// getTransactions handles GET /transactions
func (app *App) getTransactions(w http.ResponseWriter, r *http.Request) {
	// Update the stats
	app.Purge()

	if err := json.NewEncoder(w).Encode(app.Transactions.T); err != nil {
		log.Printf("Could not encode transactions: %v", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}

// addTransaction handles POST /transactions
func (app *App) addTransaction(w http.ResponseWriter, r *http.Request) {
	var body map[string]string

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Body decoding error: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, key := range transactionFields {
		if _, ok := body[key]; !ok {
			log.Printf("Body key '%s' not present in body", key)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// Try to parse float64 from the "amount" field
	f, err := strconv.ParseFloat(body["amount"], 63)
	if err != nil {
		log.Printf("Could not parse 'amount' field: %v", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Try to parse ISO 8601 timestamp
	// isoFormat := "2006-02-01T15:14:05.999Z"
	t, err := time.Parse(time.RFC3339Nano, body["timestamp"])
	if err != nil {
		log.Printf("Could not parse 'timestamp' field: %v", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	transaction := &Transaction{
		Amount:    f,
		Timestamp: t,
	}

	// Handle the incoming transaction in a atomic way
	code, _ := app.NewTx(transaction)

	// Code is -1 if its a future tx
	if code == -1 {
		log.Printf("The given timestamp (%s) is in the future", transaction.Timestamp)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Code is -2 if its a past tx (pre PurgeTime)
	if code == -2 {
		w.WriteHeader(http.StatusNoContent)
	}

	// If everything went OK return created status code
	w.WriteHeader(http.StatusCreated)
}

// deleteTransaction handles DELETE /transactions
func (app *App) deleteTransactions(w http.ResponseWriter, r *http.Request) {
	app.Transactions.Lock()
	app.Stats.Lock()

	defer app.Transactions.Unlock()
	defer app.Stats.Unlock()

	// Reset everything
	app.Transactions.T = make([]*Transaction, 0)
	app.Stats.S = newStatistic()

	// Return empty response with no content status code
	w.WriteHeader(http.StatusNoContent)
}

// getStatistics handles GET /statistics
func (app *App) getStatistics(w http.ResponseWriter, r *http.Request) {
	// Update the stats
	app.Purge()

	if err := json.NewEncoder(w).Encode(app.Stats.S); err != nil {
		log.Printf("Could not encode statistic (%v): %v", app.Stats.S, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
}
