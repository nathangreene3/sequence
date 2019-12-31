package sequence

// Comparable ...
type Comparable interface {
	Compare(c Comparable) int
}
