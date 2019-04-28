package quicksort

import (
	"math/rand"

	su "github.com/Viking2012/IntroToAlgorithms/utils"
)

// Quicksort is expected to run in O(nlogn) time, although it's
// worst case is actually O(n^2).
func Quicksort(A su.Interface) su.Interface {
	B := A

	n := rand.Intn(A.Len())
	quicksort(B, 0, n-1)

	// for left, right := 0, B.Len()-1; left < right; left, right = left+1, right-1 {
	// 	B.Swap(left, right)
	// }
	return B
}

func partition(A su.Interface, p, r int) int {
	// fmt.Printf("Paritioning at A[%d]\n", r)
	x := A.Get(r)
	// fmt.Printf("Paritioning at A[%d] and got %d\n", r, x)
	i := p
	// fmt.Println("Partitioning and setting i to ", i)

	for j := p; j < r; j++ {
		y := A.Get(j)
		// fmt.Printf("Checking if %d:%d is larger than %d:%d\n", i, x, j, y)
		if y <= x {
			// fmt.Printf("Swapped %d:%d with %d:%d\n", i, A.Get(i), j, A.Get(j))
			A.Swap(i, j)
			// fmt.Println("now have", A)
			i++
		}
	}
	A.Swap(i, r)
	return i
}

func quicksort(A su.Interface, p, r int) {
	if p < r {
		// fmt.Println("partitioning at ", p, r)
		q := partition(A, p, r)
		// fmt.Println("quicksorting at ", p, q-1)
		quicksort(A, p, q-1)
		// fmt.Println("quicksorting at ", p, q)
		quicksort(A, q, r)
	}
}
