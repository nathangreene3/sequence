package zmodn

import (
	"strconv"
	"strings"

	"github.com/nathangreene3/math"
)

// TODO
// 1. Move to math/int.
// 2. Rename Z to Int or something more fitting.

// Z is a big, base-n, signed integer. It behaves similar to big.Int, but is
// not limited to any particular base.
type Z struct {
	negative bool
	modulus  int
	value    []int
}

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// New returns a value represented as a Z.
func New(value int, modulus int) *Z {
	if value < 0 {
		return &Z{
			negative: true,
			modulus:  modulus,
			value:    math.Base(-value, modulus),
		}
	}

	return &Z{
		modulus: modulus,
		value:   math.Base(value, modulus),
	}
}

// Bin is equivalent to New(v,2).
func Bin(value int) *Z {
	return New(value, 2)
}

// Tern is equivalent to New(v,3).
func Tern(value int) *Z {
	return New(value, 3)
}

// Oct is equivalent to New(v,8).
func Oct(value int) *Z {
	return New(value, 8)
}

// Dec is equivalent to New(v,10).
func Dec(value int) *Z {
	return New(value, 10)
}

// Hex is equivalent to New(v,16).
func Hex(value int) *Z {
	return New(value, 16)
}

// One is equivalent to New(1,n).
func One(modulus int) *Z {
	return New(1, modulus)
}

// Zero is equivalent to New(0,n).
func Zero(modulus int) *Z {
	return New(0, modulus)
}

// ----------------------------------------------------------------------------
//
// ----------------------------------------------------------------------------

// Abs ...
func (x *Z) Abs() *Z {
	return &Z{
		modulus: x.modulus,
		value:   append(make([]int, 0, len(x.value)), x.value...),
	}
}

// Add returns the sum of x and y.
func Add(x, y *Z) *Z {
	return x.Copy().Add(y)
}

// Add y to x.
func (x *Z) Add(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("moduli do not match")
	}

	var (
		xLen, yLen = len(x.value), len(y.value)
		v0, k0     int // x(i) + y(i) = k0*n + v0
		v1, k1     int // v0 + k1 = k1*n + v1
	)

	if x.negative != y.negative {
		// sgn(x) != sgn(y) --> x + y = sgn(max{|x|,|y|})||x| - |y||
		switch c := x.CompareAbs(y); {
		case c < 0:
			// |x| < |y| --> x + y = sgn(y)(|y| - |x|)
			x.negative = y.negative
			for i := 0; i < xLen; i++ {
				v0, k0 = subtractMod(y.value[i], k1, n)
				v1, k1 = subtractMod(v0, x.value[i], n)
				x.value[i] = v1
				k1 += k0
			}

			for i := xLen; i < yLen; i++ {
				v1, k1 = subtractMod(y.value[i], k1, n)
				x.value = append(x.value, v1)
			}

			return x.trim()
		case 0 < c:
			// |x| > |y| --> x + y = sgn(x)(|x| - |y|)
			for i := 0; i < yLen; i++ {
				v0, k0 = subtractMod(x.value[i], k1, n)
				v1, k1 = subtractMod(v0, y.value[i], n)
				x.value[i] = v1
				k1 += k0
			}

			for i := yLen; i < xLen && k1 != 0; i++ {
				x.value[i], k1 = subtractMod(x.value[i], k1, n)
			}

			return x.trim()
		default:
			// |x| = |y| --> x + y = 0
			x.value = x.value[:0]
			x.negative = false
			return x
		}
	}

	// Changes below here may need to be applied to Subtract.

	minLen := math.MinInt(xLen, yLen)
	for i := 0; i < minLen; i++ {
		v0, k0 = addMod(x.value[i], y.value[i], n)
		v1, k1 = addMod(v0, k1, n)
		x.value[i] = v1
		k1 += k0
	}

	switch {
	case minLen < xLen:
		// |y| < |x|
		for i := minLen; i < xLen && k1 != 0; i++ {
			x.value[i], k1 = addMod(x.value[i], k1, n)
		}
	case minLen < yLen:
		// |x| < |y|
		for i := minLen; i < yLen; i++ {
			v1, k1 = addMod(y.value[i], k1, n)
			x.value = append(x.value, v1)
		}
	}

	// Any additional carry k1 appends repeatedly until k1 is empty
	for k1 != 0 {
		v1, k1 = addMod(0, k1, n)
		x.value = append(x.value, v1)
	}

	return x.trim()
}

// AddInt ...
func (x *Z) AddInt(y int) *Z {
	return x.Add(New(y, x.modulus))
}

// Compare ...
func (x *Z) Compare(y *Z) int {
	switch {
	case x.modulus != y.modulus:
		panic("unequal moduli")
	case x.negative:
		if y.negative {
			return -x.CompareAbs(y)
		}

		return -1
	default:
		if y.negative {
			return 1
		}

		return x.CompareAbs(y)
	}
}

// CompareAbs ...
func (x *Z) CompareAbs(y *Z) int {
	xLen, yLen := len(x.value), len(y.value)
	switch {
	case xLen < yLen:
		return -1
	case yLen < xLen:
		return 1
	default:
		for i := xLen - 1; 0 <= i; i-- {
			switch {
			case x.value[i] < y.value[i]:
				return -1
			case y.value[i] < x.value[i]:
				return 1
			}
		}

		return 0
	}
}

// Copy an integer.
func (x *Z) Copy() *Z {
	return &Z{
		negative: x.negative,
		modulus:  x.modulus,
		value:    append(make([]int, 0, len(x.value)), x.value...),
	}
}

// Divide y into x. Returns x.
func (x *Z) Divide(y *Z) *Z {
	// TODO: This is the lazy solution. Make it like Add.
	return x.divideInt(y.Integer())
}

// Divide y into x. Returns x. TODO
func (x *Z) divideInt(y int) *Z {
	return x.set(x.Integer() / y)
}

// Equal ...
func (x *Z) Equal(y *Z) bool {
	return x.Compare(y) == 0
}

// Integer returns the base-10 integer value.
// TODO: remove.
func (x *Z) Integer() int {
	if x.negative {
		return -math.Base10(x.value, x.modulus)
	}

	return math.Base10(x.value, x.modulus)
}

// IsEven ...
func (x *Z) IsEven() bool {
	xLen := len(x.value)
	if xLen == 0 {
		return true
	}

	if x.modulus&1 == 1 {
		// An even number of odd coefficients indicates x is even. This is only necessary for odd moduli.
		// For example, 102 (base 3) is odd because 1*3^2 + 2*1 has three quantities of odd moduli, but 112 is even because 1*3^2 + 1*3 + 2*1 has four quanties of odd moduli.
		var c byte // If c is 0, then there are an even number of odd coefficients. Otherwise (should equal 1), there are an odd number of odd coefficients.
		for i := 0; i < xLen; i++ {
			if x.value[i]&1 == 1 {
				c = (c + 1) % 2
			}
		}

		return c == 0
	}

	return x.value[0]&1 == 0
}

// IsNegative determines if an integer is negative. As a special case, zero is neither positive, nor negative.
func (x *Z) IsNegative() bool {
	return x.negative && 0 < len(x.value)
}

// IsOdd determines if an integer is odd.
func (x *Z) IsOdd() bool {
	return !x.IsEven()
}

// IsPositive determines if an integer is positive. As a special case, zero is neither positive, nor negative.
func (x *Z) IsPositive() bool {
	return !x.negative && 0 < len(x.value)
}

// IsZero determines if an integer is zero.
func (x *Z) IsZero() bool {
	return len(x.value) == 0
}

// Mod ...TODO
func (x *Z) Mod(n int) *Z {
	return x
}

// Multiply ...TODO
func (x *Z) Multiply(y *Z) *Z {
	// An attempt at the standard algorithm
	if len(x.value) < len(y.value) {
		var (
			n      = len(x.value)
			z      = x.Copy()
			v0, k0 int
			// v1, k1 int
		)

		// TODO: Finish this and make n the correct length
		for i := 0; i < n; i++ {
			v0, k0 = addMod(x.value[i], k0, x.modulus)
			z.value[i], k0 = addMod(y.value[i], v0, x.modulus)
		}

		return x
	}

	// TODO: This is the lazy solution. Make it like Add.
	x.negative = x.negative != y.negative

	one := One(x.modulus)
	for z := y.Copy().Abs(); z.IsPositive(); z.Subtract(one) {
		x.Add(x)
	}

	return x
}

// MultiplyInt ...
func (x *Z) MultiplyInt(y int) *Z {
	return x.Multiply(New(y, x.modulus))
}

// Negate ...
func (x *Z) Negate() *Z {
	x.negative = !x.negative
	return x
}

// normalize each indexed value to Z[n], where n is the modulus.
func (x *Z) normalize() *Z {
	xLen := len(x.value)
	if xLen == 0 {
		return x
	}

	var i, k int
	for ; i < xLen; i++ {
		x.value[i], k = addMod(x.value[i], k, x.modulus)
	}

	if i < xLen {
		// Ran out of value before i iterated to n
		if p := math.NextPowOfTwo(i); p < cap(x.value) {
			x.value = append(make([]int, 0, p), x.value[:i]...)
			return x
		}

		x.value = x.value[:i]
		return x
	}

	for v := 0; k != 0; {
		v, k = addMod(k, 0, x.modulus)
		x.value = append(x.value, v)
	}

	return x
}

// Pow ...TODO
func (x *Z) Pow(y *Z) *Z {
	return x
}

// PowInt ...
func (x *Z) PowInt(y int) *Z {
	return x.Pow(New(y, x.modulus))
}

// set ... TODO: Remove.
func (x *Z) set(value int) *Z {
	if value == 0 {
		if x.negative {
			x.negative = false
		}

		if 0 < len(x.value) {
			x.value = x.value[:0]
		}

		return x
	}

	if x.negative = value < 0; x.negative {
		value *= -1
	}

	var (
		xLen = len(x.value)
		i    int
	)

	for ; i < xLen && value != 0; i++ {
		x.value[i], value = addMod(value, 0, x.modulus)
	}

	if i < xLen {
		// Ran out of value before i iterated to n
		if p := math.NextPowOfTwo(i); p < cap(x.value) {
			x.value = append(make([]int, 0, p), x.value[:i]...)
			return x
		}

		x.value = x.value[:i]
		return x
	}

	for v := 0; value != 0; {
		v, value = addMod(value, 0, x.modulus)
		x.value = append(x.value, v)
	}

	return x
}

// Sgn returns the signum of x.
func (x *Z) Sgn() int {
	switch {
	case len(x.value) == 0:
		return 0
	case x.negative:
		return -1
	default:
		return 1
	}
}

// String returns a string-representation of x.
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
		panic("moduli do not match")
	}

	var (
		xLen, yLen = len(x.value), len(y.value)
		v0, k0     int
		v1, k1     int
	)

	if x.negative == y.negative {
		switch c := x.CompareAbs(y); {
		case c < 0:
			// |x| < |y| --> sgn(-x)(|y| - |x|)
			x.negative = !x.negative
			for i := 0; i < xLen; i++ {
				v0, k0 = subtractMod(y.value[i], k1, n)
				v1, k1 = subtractMod(v0, x.value[i], n)
				x.value[i] = v1
				k1 += k0
			}

			for i := xLen; i < yLen; i++ {
				v1, k1 = subtractMod(y.value[i], k1, n)
				x.value = append(x.value, v1)
			}

			return x.trim()
		case 0 < c:
			// |x| > |y| --> sgn(x)(|x| - |y|)
			for i := 0; i < yLen; i++ {
				v0, k0 = subtractMod(x.value[i], k1, n)
				v1, k1 = subtractMod(v0, y.value[i], n)
				x.value[i] = v1
				k1 += k0
			}

			for i := yLen; i < xLen && k1 != 0; i++ {
				x.value[i], k1 = subtractMod(x.value[i], k1, n)
			}

			return x.trim()
		default:
			// |x| = |y| --> 0
			x.value = x.value[:0]
			x.negative = false
			return x
		}
	}

	// Changes below here may need to be applied to Add

	minLen := math.MinInt(xLen, yLen)
	for i := 0; i < minLen; i++ {
		v0, k0 = addMod(x.value[i], y.value[i], n)
		v1, k1 = addMod(v0, k1, n)
		x.value[i] = v1
		k1 += k0
	}

	switch {
	case minLen < xLen:
		// |y| < |x|
		for i := minLen; i < xLen && k1 != 0; i++ {
			x.value[i], k1 = addMod(x.value[i], k1, n)
		}
	case minLen < yLen:
		// |x| < |y|
		for i := minLen; i < yLen; i++ {
			v1, k1 = addMod(y.value[i], k1, n)
			x.value = append(x.value, v1)
		}
	}

	// Any additional carry k1 appends repeatedly until k1 is empty
	for k1 != 0 {
		v1, k1 = addMod(0, k1, n)
		x.value = append(x.value, v1)
	}

	return x.trim()
}

// SubtractInt ...
func (x *Z) SubtractInt(y int) *Z {
	return x.Subtract(New(y, x.modulus))
}

// trim all leading zeroes and set sign if necessary.
func (x *Z) trim() *Z {
	n := len(x.value)
	for ; 0 < n && x.value[n-1] == 0; n-- {
	}

	if n == 0 {
		x.value = x.value[:n]
		x.negative = false
		return x
	}

	// TODO: Is this a bad idea?
	if p := math.NextPowOfTwo(n); p < cap(x.value) {
		x.value = append(make([]int, 0, p), x.value[:n]...)
		return x
	}

	x.value = x.value[:n]
	return x
}
