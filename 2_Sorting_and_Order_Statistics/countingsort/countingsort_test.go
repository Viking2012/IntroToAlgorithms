package countingsort

import (
	"testing"

	. "github.com/Viking2012/IntroToAlgorithms/utils"
)

type testpair struct {
	values IntSlice
	sorted IntSlice
}

var tests = []testpair{
	{IntSlice([]int{5, 2, 4, 6, 1, 3}), IntSlice([]int{1, 2, 3, 4, 5, 6})},
	{IntSlice([]int{1}), IntSlice([]int{1})},
	{IntSlice([]int{69, 10, 55, 62, 70, 92, 9, 1, 75, 38, 97, 33, 23, 90, 63, 62, 14, 55, 50, 43, 42, 40, 95, 57, 13, 21, 87, 35, 90, 43, 33, 31, 80, 33, 54, 75, 63, 15, 25, 27, 27, 79, 40, 88, 25, 7, 96, 31, 19, 38}), IntSlice([]int{1, 7, 9, 10, 13, 14, 15, 19, 21, 23, 25, 25, 27, 27, 31, 31, 33, 33, 33, 35, 38, 38, 40, 40, 42, 43, 43, 50, 54, 55, 55, 57, 62, 62, 63, 63, 69, 70, 75, 75, 79, 80, 87, 88, 90, 90, 92, 95, 96, 97})},
	{IntSlice([]int{59, 77, 36, 70, 24, 58, 29, 55, 18, 85, 62, 87, 3, 81, 43, 51, 93, 84, 75, 70, 78, 36, 57, 85, 78, 73, 42, 64, 90, 12, 63, 49, 12, 71, 26, 59, 34, 28, 78, 58, 40, 55, 61, 8, 73, 83, 46, 67, 84, 41}), IntSlice([]int{3, 8, 12, 12, 18, 24, 26, 28, 29, 34, 36, 36, 40, 41, 42, 43, 46, 49, 51, 55, 55, 57, 58, 58, 59, 59, 61, 62, 63, 64, 67, 70, 70, 71, 73, 73, 75, 77, 78, 78, 78, 81, 83, 84, 84, 85, 85, 87, 90, 93})},
	{IntSlice([]int{80, 12, 1, 57, 18, 57, 46, 3, 79, 36, 56, 44, 23, 55, 61, 89, 60, 1, 31, 91, 96, 4, 5, 79, 25, 7, 59, 38, 46, 71, 19, 5, 85, 96, 18, 11, 87, 40, 12, 81, 28, 35, 51, 1, 95, 16, 59, 5, 35, 9, 53, 12, 95, 20, 13, 12, 56, 72, 9, 4, 48, 50, 78, 64, 97, 13, 70, 25, 17, 71, 85, 50, 70, 19, 43, 13, 28, 70, 79, 54, 19, 63, 46, 17, 6, 56, 78, 36, 75, 41, 37, 62, 21, 85, 95, 87, 44, 22, 5, 44}), IntSlice([]int{1, 1, 1, 3, 4, 4, 5, 5, 5, 5, 6, 7, 9, 9, 11, 12, 12, 12, 12, 13, 13, 13, 16, 17, 17, 18, 18, 19, 19, 19, 20, 21, 22, 23, 25, 25, 28, 28, 31, 35, 35, 36, 36, 37, 38, 40, 41, 43, 44, 44, 44, 46, 46, 46, 48, 50, 50, 51, 53, 54, 55, 56, 56, 56, 57, 57, 59, 59, 60, 61, 62, 63, 64, 70, 70, 70, 71, 71, 72, 75, 78, 78, 79, 79, 79, 80, 81, 85, 85, 85, 87, 87, 89, 91, 95, 95, 95, 96, 96, 97})},
}

func TestCountingsort(t *testing.T) {
	for _, pair := range tests {
		testVal := Countingsort(pair.values)
		if !interfacesEqual(testVal, pair.sorted) {
			t.Error("\n",
				// "For     ", pair.values, "\n",
				"Expected", pair.sorted, "\n",
				"Got     ", testVal, "\n",
			)
		}
	}
}

func interfacesEqual(a Interface, b Interface) bool {
	// true! a and b are nil.
	if (a == nil) != (b == nil) {
		return true
	}

	// false! a and b are not the same length.
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		// false! element i of a and b are not the same!"
		if a.Get(i) != b.Get(i) {
			return false
		}
	}

	// true! Every element of a and b are the same.
	return true
}

var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

// var ints = []int{80, 12, 1, 57, 18, 57, 46, 3, 79, 36, 56, 44, 23, 55, 61, 89, 60, 1, 31, 91, 96, 4, 5, 79, 25, 7, 59, 38, 46, 71, 19, 5, 85, 96, 18, 11, 87, 40, 12, 81, 28, 35, 51, 1, 95, 16, 59, 5, 35, 9, 53, 12, 95, 20, 13, 12, 56, 72, 9, 4, 48, 50, 78, 64, 97, 13, 70, 25, 17, 71, 85, 50, 70, 19, 43, 13, 28, 70, 79, 54, 19, 63, 46, 17, 6, 56, 78, 36, 75, 41, 37, 62, 21, 85, 95, 87, 44, 22, 5, 44}
var result Interface

func BenchmarkCountingsort(b *testing.B) {
	var r Interface
	data := ints
	testSlice := IntSlice(data[:])
	for n := 0; n < b.N; n++ {
		r = Countingsort(testSlice)
	}

	result = r
}
