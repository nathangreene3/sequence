package sequence

// TODO: Split this into several files.

// field ...
type field []int

// newField ...
func newField(terms ...int) field {
	f := make(field, len(terms))
	copy(f, terms)
	return f
}

// compare ...
func (f field) compare(field field, iq indexQueue) int {
	if len(f) != len(field) {
		panic("dimension mismatch")
	}

	for _, index := range iq {
		x, y := f[index], field[index]
		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

// copy ...
func (f field) copy() field {
	cpy := make(field, len(f), cap(f))
	copy(cpy, f)
	return cpy
}

// start ...
type start field

// newStart ...
func newStart(terms ...int) start {
	s := make(start, len(terms))
	copy(s, terms)
	return s
}

// compare ...
func (s start) compare(start start, iq indexQueue) int {
	if len(s) != len(start) {
		panic("dimension mismatch")
	}

	for _, index := range iq {
		x, y := s[index], start[index]
		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

// copy ...
func (s start) copy() start {
	cpy := make(start, len(s), cap(s))
	copy(cpy, s)
	return cpy
}

// current ...
type current field

// newCurrent ...
func newCurrent(terms ...int) current {
	c := make(current, len(terms))
	copy(c, terms)
	return c
}

// compare ...
func (c current) compare(current current, iq indexQueue) int {
	if len(c) != len(current) {
		panic("dimension mismatch")
	}

	for _, index := range iq {
		x, y := c[index], current[index]
		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

// copy ...
func (c current) copy() current {
	cpy := make(current, len(c), cap(c))
	copy(cpy, c)
	return c
}

// end ...
type end field

// newEnd ...
func newEnd(terms ...int) end {
	e := make(end, len(terms))
	copy(e, terms)
	return e
}

// compare ...
func (e end) compare(end end, iq indexQueue) int {
	if len(e) != len(end) {
		panic("dimension mismatch")
	}

	for _, index := range iq {
		x, y := e[index], end[index]
		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

// copy ...
func (e end) copy() end {
	cpy := make(end, len(e), cap(e))
	copy(cpy, e)
	return cpy
}
