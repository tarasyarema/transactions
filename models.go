package main

import "time"

type transaction struct {
	ID        int64
	Amount    int64     `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type statistic struct {
	Sum   int64 `json:"sum"`
	Avg   int64 `json:"avg"`
	Max   int64 `json:"max"`
	Min   int64 `json:"min"`
	Count int64 `json:"count"`
}
