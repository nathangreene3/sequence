package zmodn

import (
	"github.com/nathangreene3/math"
)

// Z ...
type Z struct {
	value    []int
	modulus  int
	negative bool
}

// New ...
func New(value int, modulus int) *Z {
	if value < 0 {
		return &Z{value: math.Base(-value, modulus), modulus: modulus, negative: true}
	}

	return &Z{value: math.Base(value, modulus), modulus: modulus}
}

// Zero ...
func Zero(modulus int) *Z {
	return New(0, modulus)
}

// Abs ...
func (x *Z) Abs() *Z {
	y := x.Copy()
	y.negative = false
	return y
}

// Add ...
func (x *Z) Add(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	var bothNegative bool
	switch {
	case x.negative:
		if !y.negative {
			return y.Subtract(x.Abs())
		}

		bothNegative = true
	default:
		if y.negative {
			return x.Subtract(y.Abs())
		}
	}

	var (
		z              = New(0, n)
		xLen, yLen     = len(x.value), len(y.value)
		minLen         = math.MinInt(xLen, yLen)
		v0, v1, k0, k1 int
	)

	for i := 0; i < minLen; i++ {
		v0, k0 = addWithCarry(x.value[i], y.value[i], n)
		v1, k1 = addWithCarry(v0, k1, n)
		z.value = append(z.value, v1)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v1, k1 = addWithCarry(y.value[i], k1, n)
			z.value = append(z.value, v1)
		}
	case yLen:
		for i := minLen; i < xLen; i++ {
			v1, k1 = addWithCarry(x.value[i], k1, n)
			z.value = append(z.value, v1)
		}
	default:
		panic("")
	}

	z.negative = bothNegative
	z.clean()
	return z
}

// clean ...
func (x *Z) clean() {
	x.trim()
	x.normalize()
}

// Compare ...
func (x *Z) Compare(y *Z) int {
	var bothNegative bool
	switch {
	case x.negative:
		if !y.negative {
			return 1
		}
		bothNegative = true
	case y.negative:
		return -1
	}

	xLen, yLen := len(x.value), len(y.value)
	if xLen < yLen {
		for i := yLen - 1; xLen <= i; i-- {
			if y.value[i] != 0 {
				if bothNegative {
					return -1
				}
				return 1
			}
		}
	}

	if yLen < xLen {
		for i := xLen - 1; yLen <= i; i-- {
			if x.value[i] != 0 {
				if bothNegative {
					return 1
				}
				return -1
			}
		}
	}

	for i := math.MinInt(xLen, yLen) - 1; 0 <= i; i-- {
		switch {
		case x.value[i] < y.value[i]:
			if bothNegative {
				return 1
			}
			return -1
		case y.value[i] < x.value[i]:
			if bothNegative {
				return -1
			}
			return 1
		}
	}

	return 0
}

// Copy ...
func (x *Z) Copy() *Z {
	cpy := Z{
		value:    make([]int, len(x.value)),
		modulus:  x.modulus,
		negative: x.negative,
	}

	copy(cpy.value, x.value)
	return &cpy
}

// Integer ...
func (x *Z) Integer() int {
	n := math.Base10(x.value, x.modulus)
	if x.negative {
		return -n
	}

	return n
}

// Negate ...
func (x *Z) Negate() *Z {
	y := x.Copy()
	y.negative = !y.negative
	return y
}

// normalize ...
func (x *Z) normalize() {
	var k int
	for i, v := range x.value {
		x.value[i], k = addWithCarry(v, k, x.modulus)
	}
}

// Subtract ...
func (x *Z) Subtract(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	if 0 < x.Compare(y) {
		x, y = y, x
	}

	var bothNegative bool
	switch {
	case x.negative:
		if !y.negative {
			return y.Subtract(x.Abs())
		}

		bothNegative = true
	default:
		if y.negative {
			return x.Add(y.Abs())
		}
	}

	var (
		z              = New(0, n)
		xLen, yLen     = len(x.value), len(y.value)
		minLen         = math.MinInt(xLen, yLen)
		v0, v1, k0, k1 int
	)

	for i := 0; i < minLen; i++ {
		v0, k0 = subtractWithBorrow(x.value[i], y.value[i], n)
		v1, k1 = subtractWithBorrow(v0, k1, n)
		z.value = append(z.value, v1)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {

		}
	case yLen:
		for i := minLen; i < xLen; i++ {

		}
	}

	z.negative = bothNegative
	z.clean()
	return z
}

// trim ...
func (x *Z) trim() {
	var (
		n = len(x.value)
		c int
	)

	for i := n - 1; 0 <= i && x.value[i] == 0; i-- {
		c++
	}

	x.value = x.value[:n-c]
}
