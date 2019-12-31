package sequence

// ints ...
type ints struct {
	dims        int
	start       start
	current     current
	end         end
	format      format
	indQueue    indexQueue
	overflowed  bool
	underflowed bool
}

// newInts ...
func newInts(opts ...interface{}) ints {
	var its ints
	for _, opt := range opts {
		switch t := opt.(type) {
		case start:
			its.start = newStart(t...)
		case current:
			its.current = newCurrent(t...)
		case end:
			its.end = newEnd(t...)
		case format:
			its.dims = len(t)
			its.format = newFormat(t...)
		case indexQueue:
			its.indQueue = newOrder(t...)
		default:
			panic("invalid option")
		}
	}

	if its.format == nil {
		// TODO: Provide a default format allowing ints to be expanded in dimension.
		panic("format required")
	}

	if its.indQueue == nil {
		its.indQueue = make(indexQueue, 0, its.dims)
		for i := its.dims - 1; 0 <= i; i-- {
			its.indQueue = append(its.indQueue, i)
		}
	}

	if its.start == nil {
		its.start = make(start, 0, its.dims)
		for _, bf := range its.format {
			its.start = append(its.start, bf.min)
		}
	}

	if its.current == nil {
		its.current = current(its.start.copy())
	}

	if its.end == nil {
		its.end = make(end, 0, its.dims)
		for _, bf := range its.format {
			its.end = append(its.end, bf.max)
		}
	}

	if !its.indQueue.isValid(its.format) {
		panic("invalid order")
	}

	return its
}

// increment ...
func (its *ints) increment() {
	var (
		lenOrd = len(its.indQueue)
		carry  = 1
		count  int
	)

	for _, index := range its.indQueue {
		its.current[index], carry = its.format[index].addWithCarry(its.current[index], carry)
		if its.current[index] == its.format[index].max {
			if count++; count == lenOrd {
				its.overflowed = true
			}
		}
	}
}

// decrement ...
func (its *ints) decrement() {
	var (
		lenOrd = len(its.indQueue)
		borrow = 1
		count  int
	)

	for _, index := range its.indQueue {
		its.current[index], borrow = its.format[index].subtractWithBorrow(its.current[index], borrow)
		if its.current[index] == its.format[index].min {
			if count++; count == lenOrd {
				its.underflowed = true
			}
		}
	}
}

// add ...
func (its *ints) add(c field) {
	var carry, count int
	for _, index := range its.indQueue {
		its.current[index], carry = its.format[index].addWithCarry(its.current[index], c[index]+carry)
		if count++; count == its.dims {
			its.overflowed = true
		}
	}
}

// subtract ...
func (its *ints) subtract(c field) {
	// TODO
}
