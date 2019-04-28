package countingsort

import (
	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

// Countingsort is expected to run in O(nlogn) time.
func Countingsort(A su.Interface) su.Interface {
	B := make(su.IntSlice, A.Len())

	k := 0

	for i := 0; i < A.Len(); i++ {
		a := A.Get(i)
		// fmt.Println(i, "In setting k, checking if", a, "is larger than", k)
		if a > k {
			k = a
		}
	}
	// fmt.Println("k is", k)

	countingsort(A, B, k)

	return B
}

func countingsort(A, B su.Interface, k int) {
	C := make(su.IntSlice, k+1)

	for j := 0; j < A.Len(); j++ {
		a := A.Get(j)
		// fmt.Printf("A[%d] = %d\n", j, a)
		c := C.Get(a)
		// fmt.Printf("C[%d] = %d\n", a, c)
		C.Set(a, c+1)
	}

	// fmt.Println("intial Countingsort C:", C)

	for i := 1; i <= k; i++ {
		c := C.Get(i) + C.Get(i-1)
		C.Set(i, c)
	}
	// fmt.Println("summed Countingsort C:", C)

	for j := A.Len() - 1; j >= 0; j-- {
		a := A.Get(j)
		// fmt.Printf("A[%d] = %d\n", j, a)
		c := C.Get(a)
		// fmt.Printf("C[%d] = %d\n", a, c)
		B.Set(c-1, a)
		C.Set(a, c-1)
		// fmt.Println("Countingsort B", B)
	}
}
