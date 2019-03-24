package main

import (
	"fmt"
)

func main() {
	xs := []int{5, 2, 4, 6, 1, 3}
	insertionSort(xs)
	// fmt.Println(insertionSort(xs))
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

func insertionSort(x []int) []int {
	i := 0 // needed to collect the last index of the downward shift
	for j := range x {
		key := x[j]
		for i = j - 1; i >= 0 && key < x[i]; i-- {
			x[i+1] = x[i]
		}
		x[i+1] = key
	}
	return x
}
