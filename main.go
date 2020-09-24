package main

import (
	"fmt"
	"log"
	"net/http"
)

func transactionRouter(w http.ResponseWriter, r *http.Request) {
	addTransaction(w, r)
}

func statisticsRouter(w http.ResponseWriter, r *http.Request) {
	getStatistics(w, r)
}

func main() {
	// The http port that the API will listen to
	port := 8080

	// Define the http handlers
	http.HandleFunc("/transactions", transactionRouter)
	http.HandleFunc("/statistics", statisticsRouter)

	portStr := fmt.Sprintf(":%d", port)

	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
