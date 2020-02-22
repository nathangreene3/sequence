package zmodn

import (
	"testing"
)

func TestField(t *testing.T) {
	tests := []struct {
		n, x, y, exp int
		subtract     bool
	}{
		{
			n:   10,
			x:   4,
			y:   3,
			exp: 7,
		},
		{
			n:   10,
			x:   3,
			y:   4,
			exp: 7,
		},
		{
			n:   3,
			x:   16, //  121
			y:   8,  // + 22
			exp: 24, //  220
		},
		{
			n:   3,
			x:   8,  //   22
			y:   16, // +121
			exp: 24, //  220
		},

		// Subtract
		{
			n:        10,
			x:        4,
			y:        3,
			exp:      1,
			subtract: true,
		},
		{
			n:        10,
			x:        3,
			y:        4,
			exp:      -1,
			subtract: true,
		},
		{
			n:        3,
			x:        16,
			y:        8,
			exp:      8,
			subtract: true,
		},
		{
			n:        3,
			x:        8,
			y:        16,
			exp:      -8,
			subtract: true,
		},
	}

	for _, test := range tests {
		x, y := New(test.x, test.n), New(test.y, test.n)
		if test.subtract {
			z := x.Subtract(y)
			if rec := z.Integer(); test.exp != rec {
				t.Errorf("\nexpected (%d-%d) mod %d = %d\nreceived %d\nx = %v\ny = %v\nz = %v\n", test.x, test.y, test.n, test.exp, rec, x, y, z)
			}
		} else {
			z := x.Add(y)
			if rec := z.Integer(); test.exp != rec {
				t.Errorf("\nexpected (%d+%d) mod %d = %d\nreceived %d\nx = %v\ny = %v\nz = %v\n", test.x, test.y, test.n, test.exp, rec, x, y, z)
			}
		}
	}
}

func TestAddSubInt(t *testing.T) {
	tests := []struct {
		x, y, n, exp int
		sub          bool
	}{
		{
			x:   11,
			y:   5,
			n:   3,
			exp: 16,
		},
		{
			x:   11,
			y:   -5,
			n:   3,
			exp: 6,
		},
		{
			x:   -11,
			y:   5,
			n:   3,
			exp: -6,
		},
		{
			x:   11,
			y:   5,
			n:   3,
			exp: 16,
		},

		// Subtract
		{
			x:   11,
			y:   5,
			n:   3,
			exp: 6,
			sub: true,
		},
		{
			x:   11,
			y:   -5,
			n:   3,
			exp: 16,
			sub: true,
		},
		{
			x:   -11,
			y:   5,
			n:   3,
			exp: -16,
			sub: true,
		},
		{
			x:   11,
			y:   5,
			n:   3,
			exp: 6,
			sub: true,
		},
	}

	for _, test := range tests {
		if test.sub {
			rec := New(test.x, test.n).subtractInt(test.y).Integer()
			if test.exp != rec {
				t.Fatalf("\nexpected %d\nreceived %d\n", test.exp, rec)
			}
		} else {
			rec := New(test.x, test.n).addInt(test.y).Integer()
			if test.exp != rec {
				t.Fatalf("\nexpected %d\nreceived %d\n", test.exp, rec)
			}
		}
	}
}

// incrementor is a function that increments n according to the index queue.
type incrementor func(n int) int

// decrementor is a function that decrements n according to the index queue.
type decrementor func(n int) int

func newIncDec(base int, incOrd []int) (incrementor, decrementor) {
	if base < 2 {
		panic("")
	}

	var (
	// n = len(incOrd)
	// p        = 1
	// basePows = math.BasePows(n, base)
	)

	var (
		c int
		f = func(n int) int {
			c++
			return 0
		}

		g = func(n int) int {
			c--
			return 0
		}
	)

	return f, g
}

func TestNewIncDec(t *testing.T) {

}
