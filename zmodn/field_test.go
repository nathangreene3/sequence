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
			continue
		}

		z := x.Add(y)
		if rec := z.Integer(); test.exp != rec {
			t.Errorf("\nexpected (%d+%d) mod %d = %d\nreceived %d\nx = %v\ny = %v\nz = %v\n", test.x, test.y, test.n, test.exp, rec, x, y, z)
		}
	}
}
