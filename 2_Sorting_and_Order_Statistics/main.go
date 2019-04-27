package main

import (
	"fmt"

	"github.com/Viking2012/IntroToAlgorithms/2_Sorting_and_Order_Statistics/heapsort"

	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

func main() {
	// xs := []int{4, 3, 2, 1}
	toHeap := su.IntSlice([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	// fromHeap := toHeap

	fromHeap := heapsort.HeapSort(toHeap)

	fmt.Println(fromHeap)
}
