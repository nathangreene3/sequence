package zmodn

import (
	"testing"
)

func TestField(t *testing.T) {
	tests := []struct {
		n, x, y, exp int
	}{
		{
			n:   10,
			x:   3,
			y:   4,
			exp: 7,
		},
		{
			n:   3,
			x:   16,
			y:   8,
			exp: 24,
		},
	}

	for _, test := range tests {
		var (
			x, y = New(test.x, test.n), New(test.y, test.n)
			z    = x.Add(y)
		)

		if rec := z.Integer(); test.exp != rec {
			t.Fatalf("\nexpected (%d+%d) mod %d = %d\nreceived %d\nx = %v\ny = %v\nz = %v\n", test.x, test.y, test.n, test.exp, rec, x, y, z)
		}
	}
}
