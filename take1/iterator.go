package sequence

// Iterator ...
type Iterator interface {
	Increment(n int)
	Decrement(n int)
	Compare(iterator Iterator) int
}
