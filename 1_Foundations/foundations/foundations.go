package foundations

import (
	"fmt"
)

// Interface is a copy of the native Go sort package
type Interface interface {
	// Len is the number of elements in the collection
	Len() int

	// Less reports whether the element with index i should
	// sort before the element with index j
	Less(i, j int) bool

	// Swap swaps the elements with index i and j
	Swap(i, j int)

	// These are mine, getters of the value at an index
	// and setters of a value to an index
	Get(i int) int
	Set(i, x int)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order
// This is from the native Go sort package
type IntSlice []int

// This is from the native Go sort package
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Get returns the value of the Interface as the given index
func (p IntSlice) Get(i int) int { return p[i] }

// Set takes an index and a value and places the values at the index of the array
func (p IntSlice) Set(i, x int) { p[i] = x }

// insertionSortNative is a copy of the Go sort function
func insertionSortNative(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// native is a helper function to send an Interface to Go's native insertionSort
// and return a sorted Interface
func native(data Interface) Interface {
	A := data
	n := data.Len()
	insertionSortNative(A, 0, n)
	return A
}

// InsertionSort will sort an array
func InsertionSort(data Interface) Interface {
	i := 0                   // needed to collect the last indexs of the shift
	n := data.Len()          // get the length of the array to iterate over
	A := data                // slice to contain the sorted array
	for j := 0; j < n; j++ { // iterate over all the items in the array
		key := A.Get(j)
		for i = j - 1; (i >= 0) && (key < A.Get(i)); i-- {
			A.Set(i+1, A.Get(i)) // shift greater values the up array
		}
		A.Set(i+1, key) // shift the initializing amount to the bottom of the array
	}
	return A
}

func insertionSortOld(xs []int) (A []int) {
	A = make([]int, len(xs))
	copy(A, xs)
	for j := range A {
		key := A[j]
		i := j - 1
		for i >= 0 && key < A[i] {
			fmt.Println("Replacing:", A[i+1], "with:", A[i])
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	fmt.Println(A)
	return A
}
