package main

import (
	"encoding/json"
	"math/big"
	"net/http"
	"time"
)

func (app *App) getStatistics(w http.ResponseWriter, r *http.Request) {
	stats := Statistic{
		Sum:   big.NewFloat(0),
		Avg:   big.NewFloat(0),
		Max:   big.NewFloat(0),
		Min:   big.NewFloat(0),
		Count: len(app.Transactions),
	}

	now := time.Now()

	for i, transaction := range app.Transactions {
		// Only consider transactions in a 60s interval
		if diff := now.Sub(transaction.Timestamp); diff > 60*time.Second {
			continue
		}

		stats.Sum = stats.Sum.Add(stats.Sum, transaction.Amount)

		if i == 0 {
			stats.Max = transaction.Amount
			stats.Min = transaction.Amount
			continue
		}

		if transaction.Amount.Cmp(stats.Max) == +1 {
			stats.Max = transaction.Amount
		}

		if transaction.Amount.Cmp(stats.Min) == -1 {
			stats.Min = transaction.Amount
		}
	}

	if stats.Count > 0 {
		stats.Avg = stats.Sum.Quo(stats.Sum, big.NewFloat(float64(stats.Count)))
	}

	// Write contents and
	json.NewEncoder(w).Encode(stats)
}
