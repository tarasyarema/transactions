package main

import (
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (app *App) getTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(app.Transactions)
}

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

	z := new(big.Float)

	// Try to parse arbitrary precision big.Float from
	// the "amount" field
	f, _, err := z.Parse(body["amount"], 10)
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

	now := time.Now()

	// Check if the given transaction time is in the future
	if t.After(now) {
		log.Printf("The given timestamp (%s) is in the future", t)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Check if the transaction is older than 60 seconds
	if diff := t.Sub(now); diff < -60*time.Second {
		w.WriteHeader(http.StatusNoContent)
	}

	app.Transactions = append(app.Transactions, Transaction{
		Amount:    f,
		Timestamp: t,
	})

	// If everything went OK return created status code
	w.WriteHeader(http.StatusCreated)
}

func (app *App) deleteTransactions(w http.ResponseWriter, r *http.Request) {
	// Handle the deletion of all the transactions
	app.Transactions = make([]Transaction, 0)

	// Return empty response with no content status code
	w.WriteHeader(http.StatusNoContent)
}
