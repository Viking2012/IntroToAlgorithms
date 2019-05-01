package main

import (
	"fmt"

	"github.com/Viking2012/IntroToAlgorithms/2_Sorting_and_Order_Statistics/countingsort"
	"github.com/Viking2012/IntroToAlgorithms/2_Sorting_and_Order_Statistics/heapsort"
	"github.com/Viking2012/IntroToAlgorithms/2_Sorting_and_Order_Statistics/quicksort"
	"github.com/Viking2012/IntroToAlgorithms/2_Sorting_and_Order_Statistics/radixsort"
	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

func main() {
	// xs := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	xs := []int{80, 12, 1, 57, 18, 57, 46, 3, 79, 36, 56, 44, 23, 55, 61, 89, 60, 1, 31, 91, 96, 4, 5, 79, 25, 7, 59, 38, 46, 71, 19, 5, 85, 96, 18, 11, 87, 40, 12, 81, 28, 35, 51, 1, 95, 16, 59, 5, 35, 9, 53, 12, 95, 20, 13, 12, 56, 72, 9, 4, 48, 50, 78, 64, 97, 13, 70, 25, 17, 71, 85, 50, 70, 19, 43, 13, 28, 70, 79, 54, 19, 63, 46, 17, 6, 56, 78, 36, 75, 41, 37, 62, 21, 85, 95, 87, 44, 22, 5, 44}
	c := make([]int, len(xs))
	fmt.Println("Original", xs)

	copy(c, xs)
	toSort := su.IntSlice(c)
	// fmt.Println("ToSort  ", toSort)
	fromHeap := heapsort.HeapSort(toSort)
	fmt.Println("Heaped  ", fromHeap)

	copy(c, xs)
	toSort = su.IntSlice(c)
	// fmt.Println("ToSort  ", toSort)
	fromQuick := quicksort.Quicksort(toSort)
	fmt.Println("Quick'd ", fromQuick)

	copy(c, xs)
	toSort = su.IntSlice(c)
	// fmt.Println("ToSort  ", toSort)
	fromCounting := countingsort.Countingsort(toSort)
	fmt.Println("Counted ", fromCounting)

	copy(c, xs)
	toSort = su.IntSlice(c)
	// fmt.Println("ToSort  ", toSort)
	fromRadix := radixsort.Radixsort(toSort)
	fmt.Println("Radix   ", fromRadix)
}
