package zmodn

import (
	"github.com/nathangreene3/math"
)

// Z ...
type Z struct {
	z []int
	n int
}

// New ...
func New(value int, modulus int) Z {
	return Z{z: math.Base(value, modulus), n: modulus}
}

// Add ...
func Add(x, y Z) Z {
	n := x.n
	if n != y.n {
		panic("")
	}

	// TODO: Compare x to y

	var (
		z          = New(0, n)
		xLen, yLen = len(x.z), len(y.z)
		minLen     = math.MinInt(xLen, yLen)
	)

	var v0, v1, k0, k1 int
	for i := 0; i < minLen; i++ {
		v0, k0 = addWithCarry(x.z[i], y.z[i], n)
		v1, k1 = addWithCarry(v0, k1, n)
		z.z = append(z.z, v1)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v1, k1 = addWithCarry(y.z[i], k1, n)
			z.z = append(z.z, y.z[i])
		}
	case yLen:
		for i := minLen; i < xLen; i++ {
			v1, k1 = addWithCarry(x.z[i], k1, n)
			z.z = append(z.z, v1)
		}
	}

	return z
}

// clean ...
func (z *Z) clean() {
	var (
		n = len(z.z)
		c int
	)

	for i := n - 1; 0 <= i && z.z[i] == 0; i-- {
		c++
	}

	z.z = z.z[:n-c]
}

// Compare ...TODO
func (z *Z) Compare(x Z) int {
	return 0
}

// Int ...
func (z *Z) Int() int {
	var x int
	for i, v := range z.z {
		x += v * math.PowInt(z.n, i)
	}

	return x
}

// Subtract ...
func Subtract(x, y Z) Z {
	n := x.n
	if n != y.n {
		panic("")
	}

	z := New(0, n)
	// TODO

	return z
}
