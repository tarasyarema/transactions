package main

import (
	"time"
)

// Purge cleans the txs database from all the old
// txs that are older of the defined interval by PurgeTime const.
// Also it recomputes the the stats variable
func (a *App) Purge() {
	// As we gonna loop and change this varibale we
	// lock it during all the purging process
	a.Transactions.Lock()
	defer a.Transactions.Unlock()

	// Eventhough that we only need to change
	// the stats in some cases, its ok to lock during
	// the entire process, so noone can access to
	// partial information
	a.Stats.Lock()
	defer a.Stats.Unlock()

	now := time.Now()

	// Loop while there are invalid txs
	for {
		if len(a.Transactions.T) == 0 {
			break
		}

		pivot := a.Transactions.T[0]

		// If the last tx is in the PurgTime interval
		// then we are ok, no need to purge txs
		if now.Sub(pivot.Timestamp) < a.PurgeTime {
			break
		}

		// Update the stats
		a.Stats.S.Sum -= pivot.Amount
		a.Stats.S.Count -= 1

		// Update the average
		if a.Stats.S.Count == 0 {
			a.Stats.S.Avg = 0.0
		} else {
			a.Stats.S.Avg = a.Stats.S.Sum / float64(a.Stats.S.Count)
		}

		// Min update only when queue has elements
		if minQueueLen := len(a.Stats.S.MinQueue); minQueueLen > 0 {
			tmp := binaryDelete(a.Stats.S.MinQueue, pivot.Amount, descendingOrder)
			a.Stats.S.MinQueue = tmp

			if minQueueLen == 1 {
				a.Stats.S.Min = -MaxFloat64
			} else {
				a.Stats.S.Min = a.Stats.S.MinQueue[minQueueLen-2]
			}
		}

		// Max update, equivalent to the Min
		if maxQueueLen := len(a.Stats.S.MaxQueue); maxQueueLen > 0 {
			tmp := binaryDelete(a.Stats.S.MaxQueue, pivot.Amount, ascendingOrder)
			a.Stats.S.MaxQueue = tmp

			if maxQueueLen == 1 {
				a.Stats.S.Max = MaxFloat64
			} else {
				a.Stats.S.Max = a.Stats.S.MaxQueue[maxQueueLen-2]
			}
		}

		// Remove the old tx
		a.Transactions.T = a.Transactions.T[1:]
	}
}

// NewTx handles an incoming tx
// returns a code and the inserted index in the tx db.
//
// Return codes:
//	 0 -> Tx is valid
//	-1 -> Tx is in the future
//	-2 -> Tx is in before the PurgeTime interval limit
func (a *App) NewTx(tx *Transaction) (int, int) {
	// As we will update the stats after inserting
	// (or not) the tx, we lock now, so this function
	// is as atomic as possible
	a.Stats.Lock()
	defer a.Stats.Unlock()

	now := time.Now()

	// Handle future and old txs
	// we directly do not store them
	if tx.Timestamp.After(now) {
		return -1, 0
	}

	// Tx is outside the valid interval
	if now.Sub(tx.Timestamp) > a.PurgeTime {
		return -2, 0
	}

	// Handle insert of the tx
	// this operation is atomic respect the Transactions object
	// i.e. it locks
	i := a.Transactions.Insert(tx)

	// Handle the stats update
	a.Stats.S.Sum += tx.Amount
	a.Stats.S.Count += 1

	// Once the sum and count are updated
	// we can update the average
	a.Stats.S.Avg = a.Stats.S.Sum / float64(a.Stats.S.Count)

	// Update the max queue
	tmp, _ := binaryInsert(a.Stats.S.MaxQueue, tx.Amount, ascendingOrder)
	a.Stats.S.MaxQueue = tmp

	// Update the max
	if tx.Amount > a.Stats.S.Max {
		a.Stats.S.Max = tx.Amount
	}

	// Update the min queue
	tmp, _ = binaryInsert(a.Stats.S.MinQueue, tx.Amount, descendingOrder)
	a.Stats.S.MinQueue = tmp

	// Update the min
	if tx.Amount < a.Stats.S.Min {
		a.Stats.S.Min = tx.Amount
	}

	return 0, i
}
