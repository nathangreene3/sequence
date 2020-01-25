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

	switch {
	case x.negative:
		if !y.negative {
			return y.Subtract(x.Abs())
		}
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
	}

	z.negative = x.negative && y.negative
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

// Integer ...
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

func (x *Z) String() string {
	if len(x.value) == 0 {
		return "(0) base (" + strconv.Itoa(x.modulus) + ")"
	}

	var (
		n = len(x.value)
		b strings.Builder
	)

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

// Subtract ...
func (x *Z) Subtract(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("")
	}

	switch {
	case x.negative:
		if !y.negative {
			return y.Add(x.Abs())
		}
	default:
		if y.negative {
			return x.Add(y.Abs())
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
