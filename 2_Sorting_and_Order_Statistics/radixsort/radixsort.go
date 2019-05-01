package radixsort

import (
	"math"

	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

// Radixsort can take as little as O(d(n+k))
// when the stablesort used (Insertion here)
// runs in O(n+k) time.
func Radixsort(A su.Interface) su.Interface {
	B := A

	for i := 1; i <= largestDigit(B); i++ {
		radixSort(B, i)
	}

	return B
}

// InsertionSort will sort an array
func radixSort(A su.Interface, digit int) bool {
	i := 0                      // needed to collect the last indexs of the shift
	n := A.Len()                // get the length of the array to iterate over
	objectAtDigitFound := false // this will be set to true if there is an object in A that has a value in the digit location
	for j := 0; j < n; j++ {    // iterate over all the items in the array
		val := A.Get(j)
		key := returnDigit(A.Get(j), digit)
		for i = j - 1; i >= 0 && key < returnDigit(A.Get(i), digit); i-- {
			A.Swap(i+1, i) // shift greater values the up array
			objectAtDigitFound = true
		}
		A.Set(i+1, val) // shift the initializing amount to the bottom of the array
	}

	return objectAtDigitFound
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func returnDigit(i, n int) (d int) {
	digits := countDigits(i)

	if digits < n {
		return 0
	}

	i /= int(math.Pow10(n - 1))

	return (i % 10)
}

func largestDigit(A su.Interface) int {
	var maxDigit int
	for i := 0; i < A.Len(); i++ {
		n := countDigits(A.Get(i))
		if n > maxDigit {
			maxDigit = n
		}
	}
	return maxDigit
}
