package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func printTxs(txs *Transactions) {
	for i, t := range txs.T {
		fmt.Println(i, t.Timestamp)
	}
}

func TestBinaryInsertBase(t *testing.T) {
	arr := []float64{1, 2, 4}
	wanted := []float64{1, 2, 3, 4}

	cmp := func(a, b float64) bool {
		if a <= b {
			return true
		}

		return false
	}

	arr, i := binaryInsert(arr, 3, cmp)

	if i != 2 {
		t.Fatalf("bad insert index %d, wanted %d", i, 2)
	}

	if len(arr) != len(wanted) {
		t.Fatalf("wrong arr length %d, wanted %d", len(arr), len(wanted))
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf("insert was wrong: got arr[%d] = %f, wanted %f", i, arr[i], wanted[i])
		}
	}
}

func TestBinaryInsert(t *testing.T) {
	arr := make([]float64, 0)

	cmp := func(a, b float64) bool {
		if a <= b {
			return true
		}

		return false
	}

	n := 1000

	// Generate random tx database
	for i := 0; i < n; i++ {
		r := rand.Float64()
		tmp, _ := binaryInsert(arr, r, cmp)
		arr = tmp
	}

	if len(arr) != n {
		t.Fatalf("got wrong length %d, wanted %d", len(arr), n)
	}

	// Check that the order is correct
	// in terms of the given cmp function
	// and in terms of the expected output
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if !cmp(arr[i], arr[j]) {
				t.Fatalf("got bad function ordering: [%d] %v is after [%d] %v", i, arr[i], j, arr[j])
			}

			if arr[i] > arr[j] {
				t.Fatalf("got bad expected ordering: [%d] %v is after [%d] %v", i, arr[i], j, arr[j])
			}
		}
	}
}

func TestBinaryDeleteBase(t *testing.T) {
	initial := []float64{1, 2, 4, 5, 6, 8, 9, 10, 10, 11}
	wanted := []float64{1, 2, 4, 5, 6, 8, 9, 10, 11}

	cmp := func(a, b float64) bool {
		if a <= b {
			return true
		}

		return false
	}

	arr := binaryDelete(initial, 10, cmp)

	if len(arr) != len(wanted) {
		t.Fatalf("wrong arr length %d, wanted %d", len(arr), len(wanted))
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != wanted[i] {
			t.Fatalf("delete was wrong: got arr[%d] = %f, wanted %f", i, arr[i], wanted[i])
		}
	}
}

func TestTransactionsInsert(t *testing.T) {
	txs := &Transactions{
		T: make([]*Transaction, 0),
	}

	n := 1000

	// Generate random tx database
	for i := 0; i < n; i++ {
		tx := &Transaction{
			Amount:    rand.Float64(),
			Timestamp: randomDate(),
		}

		_ = txs.Insert(tx)
	}

	if len(txs.T) != n {
		t.Fatalf("got wrong number of txs %d, wanted %d", len(txs.T), n)
	}

	// Check that every time is correct
	for i := 0; i < len(txs.T); i++ {
		for j := i + 1; j < len(txs.T); j++ {
			if txs.T[i].Timestamp.After(txs.T[j].Timestamp) {
				t.Fatalf("got txs bad ordering: [%d] %v is after [%d] %v", i, txs.T[i].Timestamp, j, txs.T[j].Timestamp)
			}
		}
	}
}

func BenchmarkTransactionInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		txs := &Transactions{
			T: make([]*Transaction, 0),
		}

		n := 10

		// Generate random tx database
		for j := 0; j < n; i++ {
			tx := &Transaction{
				Amount:    rand.Float64(),
				Timestamp: randomDate(),
			}

			_ = txs.Insert(tx)
		}
	}
}
