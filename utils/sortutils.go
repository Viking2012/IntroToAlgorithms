package sortutils

// Interface is a copy of the native Go sort package
type Interface interface {
	// Len is the number of elements in the collection
	Len() int

	// Less reports whether the element with index i should
	// sort before the element with index j
	Less(i, j int) bool

	// Swap swaps the elements with index i and j
	Swap(i, j int)

	// These are mine, getters of the value at an index
	// and setters of a value to an index
	Get(i int) int
	Set(i, x int)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order
// This is from the native Go sort package
type IntSlice []int

// This is from the native Go sort package
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Get returns the value of the Interface as the given index
func (p IntSlice) Get(i int) int { return p[i] }

// Set takes an index and a value and places the values at the index of the array
func (p IntSlice) Set(i, x int) { p[i] = x }
