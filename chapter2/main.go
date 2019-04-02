package main

import (
	"fmt"
)

func main() {
	xs := []int{5, 2, 4, 6, 1, 3}

	insertionSorted := insertionSort(xs)
	fmt.Println("insertionSorted:", insertionSorted)
	// fmt.Println("raw values:", xs)
	L, R := mergeSortOld(xs, 0, 3, 6)
	fmt.Println("mergeSorted:", L, R)
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

func mergeSortOld(xs []int, p, q, r int) (L, R []int) {
	n1 := q - p
	n2 := r - q
	L = make([]int, n1)
	R = make([]int, n2)
	for i := range L {
		L[i] = xs[p+i]
	}
	for j := range R {
		R[j] = xs[q+j]
	}

	return L, R
}
