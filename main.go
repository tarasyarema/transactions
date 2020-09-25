package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {
	// The http port that the API will listen to
	mux := http.NewServeMux()

	app := &App{
		Transactions: Transactions{
			T: make([]*Transaction, 0),
		},
		Stats: Stats{
			S: newStatistic(),
		},
		LastID:    0,
		PurgeTime: 5 * time.Second, // PurgeTime,
	}

	// Define the http handlers
	mux.HandleFunc("/transactions", app.transactionRouter)
	mux.HandleFunc("/statistics", app.statisticsRouter)

	// Apply middleware to the requests
	afterMux := app.middleware(mux)

	portStr := fmt.Sprintf(":%d", port)
	fmt.Printf("server listening on localhost%s\n", portStr)

	if err := http.ListenAndServe(portStr, afterMux); err != nil {
		log.Fatal(err)
	}
}
