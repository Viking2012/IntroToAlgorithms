package main

import (
	"fmt"

	f "./foundations"
)

func main() {
	xs := []int{5, 2, 4, 6, 1, 3}
	iSlice := f.IntSlice(xs[:])

	insertionSorted := f.InsertionSort(iSlice)

	fmt.Println("insertionSorted:", insertionSorted)
}
