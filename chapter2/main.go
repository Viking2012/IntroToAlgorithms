package main

import (
	"fmt"
)

func main() {
	xs := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertionSort(xs))
}

func insertionSortOld(x []int) []int {
	for j := range x {
		key := x[j]
		i := j - 1
		for i >= 0 && key < x[i] {
			fmt.Println("Replacing:", x[i+1], "with:", x[i])
			x[i+1] = x[i]
			i = i - 1
		}
		x[i+1] = key
	}
	fmt.Println(x)
	return x
}

// insertionSort will sort an array
func insertionSort(x []int) []int {
	i := 0             // needed to collect the last index of the shift
	for j := range x { // iterate over all the items in the array
		key := x[j] // this is the "topmost" item in the stack
		for i = j - 1; i >= 0 && key < x[i]; i-- {
			x[i+1] = x[i] // shift greater values the up array
		}
		x[i+1] = key // shift the initializing amount to the bottom of the array
	}
	return x
}
