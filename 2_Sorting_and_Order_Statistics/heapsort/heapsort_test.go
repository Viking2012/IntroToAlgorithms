package heapsort

import (
	"testing"
)

func TestHeapsort(t *testing.T) {
	toHeap := []int{1, 10, 7, 14, 16, 4, 9, 3, 8, 2}
	fromHeap := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}

	testHeap := heapsort(toHeap)

	// r := slicesEqual(testHeap[:], fromHeap[:])

	if !slicesEqual(testHeap[:], fromHeap[:]) {
		t.Error(
			"\nFrom heapsort:\t", toHeap,
			"\nexpected\t\t", fromHeap,
			"\ngot\t\t\t\t", testHeap,
		)
	}
}

func slicesEqual(a []int, b []int) bool {
	// true! a and b are nil.
	if (a == nil) != (b == nil) {
		return true
	}

	// false! a and b are not the same length.
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		// false! element i of a and b are not the same!"
		if a[i] != b[i] {
			return false
		}
	}

	// true! Every element of a and b are the same.
	return true
}
