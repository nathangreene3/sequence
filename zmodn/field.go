package zmodn

import (
	"strconv"
	"strings"

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

// One is equivalent to New(1,n).
func One(modulus int) *Z {
	return New(1, modulus)
}

// Zero is equivalent to New(0,n).
func Zero(modulus int) *Z {
	return New(0, modulus)
}

// Abs ...
func (x *Z) Abs() *Z {
	y := x.Copy()
	y.negative = false
	return y
}

// Add y to x. Returns x.
func (x *Z) Add(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	switch {
	case x.negative:
		if !y.negative {
			return x.Subtract(y)
		}
	default:
		if y.negative {
			return x.Subtract(y)
		}
	}

	var (
		xLen, yLen     = len(x.value), len(y.value)
		minLen         = math.MinInt(xLen, yLen)
		v0, v1, k0, k1 int
	)

	for i := 0; i < minLen; i++ {
		v0, k0 = addWithCarry(x.value[i], y.value[i], n)
		v1, k1 = addWithCarry(v0, k1, n)
		x.value[i] = v1
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v1, k1 = addWithCarry(y.value[i], k1, n)
			x.value[i] = v1
		}
	case yLen:
		for i := minLen; i < xLen; i++ {
			v1, k1 = addWithCarry(x.value[i], k1, n)
			x.value[i] = v1
		}
	}

	// x.negative = x.negative && y.negative
	return x.trim().normalize()
}

// Add ...
func Add(x, y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	switch {
	case x.negative:
		if !y.negative {
			return Subtract(y, x.Abs())
		}
	default:
		if y.negative {
			return Subtract(x, y.Abs())
		}
	}

	/*
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
		}

		z.negative = x.negative && y.negative
		return z.clean()
	*/

	return x.Copy().Add(y)
}

// clean calls trim and normalize.
func (x *Z) clean() *Z {
	return x.trim().normalize()
}

// Compare ...
func (x *Z) Compare(y *Z) int {
	// The modulus makes comparison difficult, so compare on base-10 representation
	xInt, yInt := x.Integer(), y.Integer()
	switch {
	case xInt < yInt:
		return -1
	case yInt < xInt:
		return 1
	default:
		return 0
	}
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

// Integer returns the base-10 integer value.
func (x *Z) Integer() int {
	n := math.Base10(x.value, x.modulus)
	if x.negative {
		n *= -1
	}

	return n
}

// IsEven ...
func (x *Z) IsEven() bool {
	return len(x.value) == 0 || x.value[0]%2 == 0
}

// IsNegative ...
func (x *Z) IsNegative() bool {
	return x.negative
}

// IsOdd ...
func (x *Z) IsOdd() bool {
	return len(x.value) != 0 && x.value[0]%2 != 0
}

// IsPositive ...
func (x *Z) IsPositive() bool {
	return !x.negative
}

// IsZero ...
func (x *Z) IsZero() bool {
	return x.Integer() == 0
}

// Mulitply ...
func (x *Z) Mulitply(y *Z) *Z {
	var (
		yInt       = y.Integer()
		isNegative = yInt < 0
	)

	if isNegative {
		yInt *= -1
	}

	z := x.Copy()
	for ; 1 < yInt; yInt-- {
		z = Add(z, x)
	}

	if isNegative {
		z = z.Negate()
	}

	return z
}

// Negate ...
func (x *Z) Negate() *Z {
	y := x.Copy()
	y.negative = !y.negative
	return y
}

// normalize each indexed value to Z[n], where n is the modulus.
func (x *Z) normalize() *Z {
	var k int
	for i, v := range x.value {
		x.value[i], k = addWithCarry(v, k, x.modulus)
	}

	return x
}

func (x *Z) String() string {
	n := len(x.value)
	if n == 0 {
		return "(0) base (" + strconv.Itoa(x.modulus) + ")"
	}

	var b strings.Builder
	if x.negative {
		b.WriteByte('-')
	}

	b.WriteString("(" + strconv.Itoa(x.value[n-1]))
	for i := n - 2; 0 <= i; i-- {
		b.WriteString("," + strconv.Itoa(x.value[i]))
	}

	b.WriteString(") (base " + strconv.Itoa(x.modulus) + ")")
	return b.String()
}

// Subtract y from x.
func (x *Z) Subtract(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	switch {
	case x.negative:
		if !y.negative {
			// return Add(y, x.Abs())
		}
	default:
		if y.negative {
			// return Add(x, y.Abs())
		}
	}

	var isNegative bool
	if x.Compare(y) < 0 {
		x, y = y, x
		isNegative = true
	}

	var (
		xLen, yLen     = len(x.value), len(y.value)
		minLen         = math.MinInt(xLen, yLen)
		v0, v1, k0, k1 int
	)

	for i := 0; i < minLen; i++ {
		v0, k0 = subtractWithBorrow(x.value[i], k1, n)
		v1, k1 = subtractWithBorrow(v0, y.value[i], n)
		x.value[i] = v1
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v1, k1 = subtractWithBorrow(y.value[i], k1, n)
			x.value[i] = v1
		}
	case yLen:
		// Should always be this case...
		for i := minLen; i < xLen; i++ {
			v1, k1 = subtractWithBorrow(x.value[i], k1, n)
			x.value[i] = v1
		}
	}

	x.negative = isNegative || (x.negative && y.negative)
	return x.clean()
}

// Subtract ...
func Subtract(x, y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	switch {
	case x.negative:
		if !y.negative {
			return Add(y, x.Abs())
		}
	default:
		if y.negative {
			return Add(x, y.Abs())
		}
	}

	var isNegative bool
	if x.Compare(y) < 0 {
		x, y = y, x
		isNegative = true
	}

	var (
		z              = New(0, n)
		xLen, yLen     = len(x.value), len(y.value)
		minLen         = math.MinInt(xLen, yLen)
		v0, v1, k0, k1 int
	)

	for i := 0; i < minLen; i++ {
		v0, k0 = subtractWithBorrow(x.value[i], k1, n)
		v1, k1 = subtractWithBorrow(v0, y.value[i], n)
		z.value = append(z.value, v1)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v1, k1 = subtractWithBorrow(y.value[i], k1, n)
			z.value = append(z.value, v1)
		}
	case yLen:
		// Should always be this case...
		for i := minLen; i < xLen; i++ {
			v1, k1 = subtractWithBorrow(x.value[i], k1, n)
			z.value = append(z.value, v1)
		}
	}

	z.negative = isNegative || (x.negative && y.negative)
	z.clean()
	return z
}

// trim all leading zeroes.
func (x *Z) trim() *Z {
	var (
		n = len(x.value)
		c int
	)

	for i := n - 1; 0 <= i && x.value[i] == 0; i-- {
		c++
	}

	x.value = x.value[:n-c]
	return x
}
