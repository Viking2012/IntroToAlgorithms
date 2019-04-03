package heapsort

// heapsort is expected to run in O(n lgn) time
func heapsort(xs []int) (A []int) {
	A = make([]int, len(xs))
	copy(A, xs)
	return A
}
