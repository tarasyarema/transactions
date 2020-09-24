package main

import (
	"math/big"
	"time"
)

// App defines the general application struct
type App struct {
	Transactions []Transaction
}

// Transaction defines the default trnasaction struct
type Transaction struct {
	ID        int64      `json:"-"`
	Amount    *big.Float `json:"amount"`
	Timestamp time.Time  `json:"timestamp"`
}

// transactionsFields return the transaction struct
// JSON fields for valid JSON checking purposes
var transactionFields = [2]string{
	"amount",
	"timestamp",
}

// Statistic defines the default statistic struct
type Statistic struct {
	Sum   *big.Float `json:"sum"`
	Avg   *big.Float `json:"avg"`
	Max   *big.Float `json:"max"`
	Min   *big.Float `json:"min"`
	Count int        `json:"count"`
}
