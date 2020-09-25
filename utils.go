package main

import (
	log "github.com/sirupsen/logrus"
)

// ascendingOrder is a utility function for the
// Max/Min queues of the stats
func ascendingOrder(a, b float64) bool {
	if a <= b {
		return true
	}

	return false
}

// descendingOrder is a utility function for the
// Max/Min queues of the stats
func descendingOrder(a, b float64) bool {
	if a >= b {
		return true
	}

	return false
}

// insertTransaction inserts the a slice of txs into s in the position k
// stolen from https://github.com/golang/go/wiki/SliceTricks#insertvector
// for the most "optimal" way to do it without packages
func insertTransaction(s []*Transaction, k int, vs ...*Transaction) []*Transaction {
	if n := len(s) + len(vs); n <= cap(s) {
		s2 := s[:n]
		copy(s2[k+len(vs):], s[k:])
		copy(s2[k:], vs)
		return s2
	}

	s2 := make([]*Transaction, len(s)+len(vs))

	copy(s2, s[:k])
	copy(s2[k:], vs)
	copy(s2[k+len(vs):], s[k:])

	return s2
}

// insertFloat64 does the same as insertTransaction but for float64 slice
// the need of a duplicate is basically that Go does not support generics (for the moment)
func insertFloat64(s []float64, k int, vs ...float64) []float64 {
	if n := len(s) + len(vs); n <= cap(s) {
		s2 := s[:n]
		copy(s2[k+len(vs):], s[k:])
		copy(s2[k:], vs)
		return s2
	}

	s2 := make([]float64, len(s)+len(vs))

	copy(s2, s[:k])
	copy(s2[k:], vs)
	copy(s2[k+len(vs):], s[k:])

	return s2
}

// binaryInsert make a binary insertion of the element e
// into the ordered slice arr with the given order funcion cmp.
//
// The function cmp(a, b) should return true if the
// relation a <= b holds, false instead.
//
// Example:
//   Given arr = [1, 2, 4], e = 3 and cmp definied by the normal <=,
//	 then binaryInsert would make
//		arr = [1, 2, 3, 4], and return 2.
func binaryInsert(arr []float64, e float64, cmp func(a, b float64) bool) ([]float64, int) {
	l := len(arr)
	if l == 0 {
		arr = append(arr, e)
		return arr, 0
	}

	// Insert e as the new first element
	if cmp(e, arr[0]) {
		arr = insertFloat64(arr, 0, e)
		return arr, 0
	}

	// Insert e as the last element
	if cmp(arr[l-1], e) {
		arr = append(arr, e)
		return arr, l - 1
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

		if cmp(arr[mid], e) {
			low = mid
		} else {
			high = mid
		}

		// If we found a space of length one
		// we already done
		if high-low <= 1 {
			if cmp(arr[low], e) {
				pos = high
			} else {
				pos = low
			}

			break
		}
	}

	arr = insertFloat64(arr, pos, e)
	return arr, pos
}

// binaryDelete make a binary deletion of the element e
// into the ordered slice arr with the given order funcion cmp.
//
// The function cmp(a, b) should return true if the
// relation a <= b holds, false instead.
//
// Example:
//   Given arr = [1, 2, 4], e = 2 and cmp definied by the normal <=,
//	 then binaryDelete would make
//		arr = [1, 4]
func binaryDelete(arr []float64, e float64, cmp func(a, b float64) bool) []float64 {
	l := len(arr)
	if l == 0 {
		return arr
	}

	low := 0
	high := l - 1

	var (
		mid int
		pos int
	)

	for {
		// Check if we need to break
		if high-low <= 1 {
			if arr[high] == e {
				pos = high
			} else if arr[low] == e {
				pos = low
			} else {
				// Should not happen, but anyway
				// bruteforce in this situation
				found := false

				for i, x := range arr {
					if x == e {
						found = true
						pos = i
						break
					}
				}

				if !found {
					log.Printf("could not find element %v in the txs db", e)
					return arr
				}
			}

			break
		}

		// The integer division floors the result
		mid = low + (high-low)/2

		// Check if we already found the elements position
		if arr[mid] == e {
			pos = mid
			break
		}

		if cmp(arr[mid], e) {
			low = mid
		} else {
			high = mid
		}
	}

	arr = arr[:pos+copy(arr[pos:], arr[pos+1:])]
	return arr
}
