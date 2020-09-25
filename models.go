package main

import (
	"sync"
	"time"
)

// MaxFloat64 is basically what its name says xd
const MaxFloat64 = 1.797693134862315708145274237317043567981e+308

// PurgeTime is the default purge interval
const PurgeTime = 60 * time.Second

// App defines the general application struct
type App struct {
	// The db with all the txs
	Transactions Transactions

	// Current stats
	Stats Stats

	// The first tx to expire
	LastID int

	// Defines the Purge interval
	PurgeTime time.Duration
}

// Transaction defines the default trnasaction struct
type Transaction struct {
	// Amount defines the amount sent in the tx
	Amount float64 `json:"amount"`

	// Timestamp defines the time the tx was created
	Timestamp time.Time `json:"timestamp"`
}

// Transactions define a transactions struct with a lock
type Transactions struct {
	sync.Mutex
	T []*Transaction
}

// transactionsFields return the transaction struct
// JSON fields for valid JSON checking purposes
var transactionFields = [2]string{
	"amount",
	"timestamp",
}

// Statistic defines the default statistic struct
type Statistic struct {
	Sum float64 `json:"sum"`
	Avg float64 `json:"avg"`

	Max      float64   `json:"max"`
	MaxQueue []float64 `json:"max_queue"`

	Min      float64   `json:"min"`
	MinQueue []float64 `json:"min_queue"`

	Count int64 `json:"count"`
}

// Stats define a statistics struct with a lock
type Stats struct {
	sync.Mutex
	S *Statistic
}
