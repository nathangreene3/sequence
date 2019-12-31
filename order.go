package sequence

// indexQueue ... TODO: Rename to indexQueue.
type indexQueue []int

// newIndexQueue ...
func newOrder(ranks ...int) indexQueue {
	cpy := make(indexQueue, len(ranks))
	copy(cpy, ranks)
	return cpy
}

// copy ...
func (iq indexQueue) copy() indexQueue {
	cpy := make(indexQueue, len(iq), cap(iq))
	copy(cpy, iq)
	return cpy
}

func (iq indexQueue) isValid(f format) bool {
	n := len(f)
	m := make(map[int]struct{})
	for _, index := range iq {
		if index < 0 || n <= index {
			return false
		}

		v, ok := m[index]
		if ok {
			return false
		}

		m[index] = v
	}

	return true
}
