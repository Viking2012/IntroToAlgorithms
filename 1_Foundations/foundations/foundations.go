package foundations

import (
	"fmt"
)

func main() {
	xs := []int{5, 2, 4, 6, 1, 3}

	insertionSorted := insertionSort(xs)
	fmt.Println("insertionSorted:", insertionSorted)
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

// insertionSort will sort an array
func insertionSort(xs []int) (A []int) {
	i := 0                   // needed to collect the last indexs of the shift
	A = make([]int, len(xs)) // slice to contain the sorted array
	copy(A, xs)              // fill the slice with the raw values
	for j := range A {       // iterate over all the items in the array
		key := A[j] // this is the "topmost" item in the stack
		for i = j - 1; i >= 0 && key < A[i]; i-- {
			A[i+1] = A[i] // shift greater values the up array
		}
		A[i+1] = key // shift the initializing amount to the bottom of the array
	}
	return A
}

func native(data Interface) Interface {
	A := data
	n := data.Len()
	insertionSortNative(A, 0, n)
	return A
}

func insertionSortNative(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// Interface is a copy of the native Go sort package
type Interface interface {
	// Len is the number of elements in the collection
	Len() int

	// Less reports whether the element with index i should
	// sort before the element with index j
	Less(i, j int) bool

	// Swap swaps the elements with index i and j
	Swap(i, j int)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
