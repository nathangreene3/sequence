package zmodn

import (
	"testing"
)

func TestFieldAddSubtract(t *testing.T) {
	/*
		tests := []struct {
			n, x, y, exp int
			subtract     bool
		}{
			// Adding positives
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

			// Adding negatives
			{
				n:   10,
				x:   37,
				y:   -100,
				exp: -63,
			},
			{
				n:   3,
				x:   -2,
				y:   2,
				exp: 0,
			},
			{
				n:   3,
				x:   2,  //  2
				y:   -2, // -2
				exp: 0,  //  0
			},
			{
				n:   3,
				x:   -8, //  -22
				y:   16, // +121
				exp: 8,  //   22
			},
			{
				n:   3,
				x:   8,   //   22
				y:   -16, // +121
				exp: -8,  //  -22
			},

			// Subtracting
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
				x:        16, // 121
				y:        8,  // -22
				exp:      8,  //  22
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
	*/

	var (
		n0, n1 = 2, 16
		a, b   = -16, 16
	)

	for n := n0; n <= n1; n++ {
		for x := a; x <= b; x++ {
			for y := a; y <= b; y++ {
				if z := New(x, n).Add(New(y, n)).Integer(); x+y != z {
					t.Errorf("\nexpected %d+%d (n = %d) = %d\nreceived %d\n", x, y, n, x+y, z)
				}

				if z := New(x, n).Subtract(New(y, n)).Integer(); x-y != z {
					t.Errorf("\nexpected %d-%d (n = %d) = %d\nreceived %d\n", x, y, n, x-y, z)
				}
			}
		}
	}
}

func TestEvenOdd(t *testing.T) {
	for n := 2; n <= 16; n++ {
		for x := -16; x <= 16; x++ {
			if exp, rec := x&1 == 0, New(x, n).IsEven(); exp != rec {
				t.Fatalf("\n   given %d and base %d\nexpected %t\nreceived %t\n", x, n, exp, rec)
			}
		}
	}
}

func BenchmarkField(b *testing.B) {
	// TODO
}
