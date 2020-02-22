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

func Bin(value int) *Z {
	return New(value, 2)
}

func Oct(value int) *Z {
	return New(value, 8)
}

func Dec(value int) *Z {
	return New(value, 10)
}

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

// Abs ...
func (x *Z) Abs() *Z {
	y := x.Copy()
	y.negative = false
	return y
}

func (x *Z) Add(y *Z) *Z {
	return x.addInt(y.Integer())
}

func (x *Z) addInt(y int) *Z {
	return x.set(x.Integer() + y)
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

func (x *Z) divideInt(y int) *Z {
	return x.set(x.Integer() / y)
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
	return x.multiplyInt(y.Integer())
}

func (x *Z) multiplyInt(y int) *Z {
	return x.set(x.Integer() * y)
}

// Negate ...
func (x *Z) Negate() *Z {
	x.negative = !x.negative
	return x
}

// normalize each indexed value to Z[n], where n is the modulus.
func (x *Z) normalize() *Z {
	var (
		k int
		n = len(x.value)
	)

	for i := 0; i < n; i++ {
		x.value[i], k = addWithCarry(x.value[i], k, x.modulus)
	}

	return x
}

func (x *Z) set(value int) *Z {
	if x.negative = value < 0; x.negative {
		value *= -1
	}

	x.value = math.Base(value, x.modulus)
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

func (x *Z) Subtract(y *Z) *Z {
	return x.subtractInt(y.Integer())
}

func (x *Z) subtractInt(y int) *Z {
	return x.set(x.Integer() - y)
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
