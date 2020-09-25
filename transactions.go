package main

// Insert inserts a given tx into the Transactions structure
// with the correct time ordering
func (t *Transactions) Insert(tx *Transaction) int {
	t.Lock()
	defer t.Unlock()

	l := len(t.T)

	// Trivial case O(1)
	if l == 0 {
		t.T = append(t.T, tx)
		return 0
	}

	// Given TX becomes the oldest O(n)
	if tx.Timestamp.Before(t.T[0].Timestamp) {
		t.T = insertTransaction(t.T, 0, tx)
		return 0
	}

	// Given TX becomes the newest O(1)
	if tx.Timestamp.After(t.T[l-1].Timestamp) {
		t.T = append(t.T, tx)
		return l - 1
	}

	// If not one of the trivial cases
	// we follow with a binary search algorithm

	low := 0
	high := l - 1

	var (
		mid int
		pos int
	)

	for {
		// The integer division floors the result
		mid = low + (high-low)/2

		if tx.Timestamp.After(t.T[mid].Timestamp) {
			low = mid
		} else {
			high = mid
		}

		// If we found a space of length one
		// we already done
		if high-low <= 1 {
			if tx.Timestamp.After(t.T[low].Timestamp) {
				pos = high
			} else {
				pos = low
			}

			break
		}
	}

	t.T = insertTransaction(t.T, pos, tx)
	return pos
}
