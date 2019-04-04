package foundations

import (
	"fmt"

	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

// insertionSortNative is a copy of the Go sort function
func insertionSortNative(data su.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// native is a helper function to send an Interface to Go's native insertionSort
// and return a sorted Interface
func native(data su.Interface) su.Interface {
	A := data
	n := data.Len()
	insertionSortNative(A, 0, n)
	return A
}

// InsertionSort will sort an array
func InsertionSort(data su.Interface) su.Interface {
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
