package sequence

// order ...
type order []int

// field ...
type field []int

// start ...
type start field

// current ...
type current field

// end ...
type end field

// baseFmt ...
type baseFmt struct {
	min, max int
}

// format ...
type format []baseFmt

// ints ...
type ints struct {
	dims       int
	start      start
	current    current
	end        end
	format     format
	order      order
	overflowed bool
}

// newBaseFmt ...
func newBaseFmt(min, max int) baseFmt {
	if max < min || min < 0 {
		panic("")
	}

	return baseFmt{min: min, max: max}
}

func newInts(opts ...interface{}) ints {
	var its ints
	for _, opt := range opts {
		switch t := opt.(type) {
		case start:
			its.start = t.copy()
		case current:
			its.current = t.copy()
		case end:
			its.end = t.copy()
		case format:
			its.dims = len(t)
			its.format = t.copy()
		case order:
			its.order = t.copy()
		default:
			panic("invalid option")
		}
	}

	if its.format == nil {
		panic("")
	}

	if its.order == nil {
		its.order = make(order, 0, its.dims)
		for i := its.dims - 1; 0 <= i; i-- {
			its.order = append(its.order, i)
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

	return its
}

// addWithCarry returns (a+b)mod(max+1) + min.
func addWithCarry(a, b int, bf baseFmt) (int, int) {
	var (
		c     = a + b
		carry int
	)

	for ; bf.max < c; carry++ {
		c -= bf.max + 1
	}

	return c, carry
}

func subtractWithBorrow(a, b int, bf baseFmt) (int, int) {
	// return addWithCarry(bf.max-bf.min-a, a, bf)
	return addWithCarry(bf.max-bf.min-a, b, bf)
}

// inc ...
func (its ints) increment() {
	var (
		carry = 1
		count int
	)

	for _, index := range its.order {
		its.current[index], carry = addWithCarry(its.current[index], carry, its.format[index])
		if count++; count == its.dims {
			its.overflowed = true
		}
	}
}

func (its ints) decrement() {
	var (
		borrow = 1
		count  int
	)

	for _, index := range its.order {
		its.current[index], borrow = subtractWithBorrow(its.current[index], borrow, its.format[index])
		if count++; count == its.dims {
			its.overflowed = true
		}
	}
}

func (its ints) add(c field) {
	var carry, count int
	for _, index := range its.order {
		its.current[index], carry = addWithCarry(its.current[index], c[index]+carry, its.format[index])
		if count++; count == its.dims {
			its.overflowed = true
		}
	}
}

func (its ints) subtract(c field) {

}

func (f field) copy() field {
	cpy := make(field, len(f), cap(f))
	copy(cpy, f)
	return cpy
}

func (s start) copy() start {
	cpy := make(start, len(s), cap(s))
	copy(cpy, s)
	return cpy
}

func (c current) copy() current {
	cpy := make(current, len(c), cap(c))
	copy(cpy, c)
	return c
}

func (e end) copy() end {
	cpy := make(end, len(e), cap(e))
	copy(cpy, e)
	return cpy
}

func (ord order) copy() order {
	cpy := make(order, len(ord), cap(ord))
	copy(cpy, ord)
	return cpy
}

func (f format) copy() format {
	cpy := make(format, len(f), cap(f))
	copy(cpy, f)
	return cpy
}

// compare ...
func (f field) compare(field field, order order) int {
	if len(f) != len(field) {
		panic("dimension mismatch")
	}

	for _, index := range order {
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
