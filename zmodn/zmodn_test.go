package zmodn

import (
	"fmt"
	"testing"
)

func TestAddSubtract(t *testing.T) {
	tests := []struct {
		x, y, n, answer, carry, borrow int
		subtract                       bool
	}{
		// 0 < x
		{
			x:      5,
			y:      4,
			n:      6,
			answer: 3,
			carry:  1,
		},
		{
			x:      5,
			y:      -4,
			n:      6,
			answer: 1,
			carry:  0,
		},
		{
			x:        5,
			y:        4,
			n:        6,
			answer:   1,
			borrow:   0,
			subtract: true,
		},
		{
			x:        5,
			y:        -4,
			n:        6,
			answer:   3,
			borrow:   -1,
			subtract: true,
		},

		// x < 0
		{
			x:      -5,
			y:      4,
			n:      6,
			answer: 5,
			carry:  -1,
		},
		{
			x:      -5,
			y:      -4,
			n:      6,
			answer: 3,
			carry:  -2,
		},
		{
			x:        -5,
			y:        4,
			n:        6,
			answer:   3,
			borrow:   2,
			subtract: true,
		},
		{
			x:        -5,
			y:        -4,
			n:        6,
			answer:   5,
			borrow:   1,
			subtract: true,
		},
	}

	for _, test := range tests {
		if test.subtract {
			answer, borrow := subtractMod(test.x, test.y, test.n)
			if test.answer != answer || test.borrow != borrow {
				t.Fatalf("\n(%d - %d) mod %d\nexpected (%d,%d)\nreceived (%d,%d)\n", test.x, test.y, test.n, test.answer, test.borrow, answer, borrow)
			}
		} else {
			answer, carry := addMod(test.x, test.y, test.n)
			if test.answer != answer || test.carry != carry {
				t.Fatalf("\n(%d + %d) mod %d\nexpected (%d,%d)\nreceived (%d,%d)\n", test.x, test.y, test.n, test.answer, test.carry, answer, carry)
			}
		}
	}
}

func TestEuclidsCoeffs(t *testing.T) {
	tests := []struct {
		x, n, expK, expR int
	}{
		// -----------------
		// Cases where 0 < n
		// -----------------
		{
			// -6 = -2*3 + 0
			x:    -6,
			n:    3,
			expK: -2,
			expR: 0,
		},
		{
			// -5 = -2*3 + 1
			x:    -5,
			n:    3,
			expK: -2,
			expR: 1,
		},
		{
			// -4 = -2*3 + 2
			x:    -4,
			n:    3,
			expK: -2,
			expR: 2,
		},
		{
			// -3 = -1*3 + 0
			x:    -3,
			n:    3,
			expK: -1,
			expR: 0,
		},
		{
			// -2 = -1*3 + 1
			x:    -2,
			n:    3,
			expK: -1,
			expR: 1,
		},
		{
			// -1 = -1*3 + 2
			x:    -1,
			n:    3,
			expK: -1,
			expR: 2,
		},
		{
			// 0 = 0*3 + 0
			x:    0,
			n:    3,
			expK: 0,
			expR: 0,
		},
		{
			// 1 = 0*3 + 1
			x:    1,
			n:    3,
			expK: 0,
			expR: 1,
		},
		{
			// 2 = 0*3 + 2
			x:    2,
			n:    3,
			expK: 0,
			expR: 2,
		},
		{
			// 3 = 1*3 + 0
			x:    3,
			n:    3,
			expK: 1,
			expR: 0,
		},
		{
			// 4 = 1*3 + 1
			x:    4,
			n:    3,
			expK: 1,
			expR: 1,
		},
		{
			// 5 = 1*3 + 2
			x:    5,
			n:    3,
			expK: 1,
			expR: 2,
		},
		{
			// 6 = 2*3 + 0
			x:    6,
			n:    3,
			expK: 2,
			expR: 0,
		},

		// -----------------
		// Cases where n < 0
		// -----------------
		{
			// -6 = 2*-3 - 0
			x:    -6,
			n:    -3,
			expK: 2,
			expR: 0,
		},
		{
			// -5 = 1*-3 - 2
			x:    -5,
			n:    -3,
			expK: 1,
			expR: -2,
		},
		{
			// -4 = 1*-3 - 1
			x:    -4,
			n:    -3,
			expK: 1,
			expR: -1,
		},
		{
			// -3 = 1*-3 + 0
			x:    -3,
			n:    -3,
			expK: 1,
			expR: 0,
		},
		{
			// -2 = 0*-3 - 2
			x:    -2,
			n:    -3,
			expK: 0,
			expR: -2,
		},
		{
			// -1 = 0*-3 - 1
			x:    -1,
			n:    -3,
			expK: 0,
			expR: -1,
		},
		{
			// 0 = 0*-3 + 0
			x:    0,
			n:    -3,
			expK: 0,
			expR: 0,
		},
		{
			// 1 = -1*-3 - 2
			x:    1,
			n:    -3,
			expK: -1,
			expR: -2,
		},
		{
			// 2 = -1*-3 - 1
			x:    2,
			n:    -3,
			expK: -1,
			expR: -1,
		},
		{
			// 3 = -1*-3 + 0
			x:    3,
			n:    -3,
			expK: -1,
			expR: 0,
		},
		{
			// 4 = -2*-3 - 2
			x:    4,
			n:    -3,
			expK: -2,
			expR: -2,
		},
		{
			// 5 = -2*-3 -1
			x:    5,
			n:    -3,
			expK: -2,
			expR: -1,
		},
		{
			// 6 = -2*-3 + 0
			x:    6,
			n:    -3,
			expK: -2,
			expR: 0,
		},
	}

	for _, test := range tests {
		k, r := EuclidFloor(test.x, test.n)
		if test.expK != k || test.expR != r {
			t.Fatalf("\nexpected euclid(%d,%d) => (%d,%d)\nreceived (%d,%d)\n", test.x, test.n, test.expK, test.expR, k, r)
		}
	}
}

func TestEuclid2(t *testing.T) {
	// x, n := 7, -3
	// t.Error(euclidFloor(x, n))
	// t.Error(euclidTrunc(x, n))
	// t.Error(euclid(x, n))
}

func BenchmarkEuclidFloor(b *testing.B) {
	var a, n int = 1 << 62, 3
	for i := 0; i < b.N; i++ {
		benchmarkEuclidFloor(b, a, n)
	}

	for i := 0; i < b.N; i++ {
		benchmarkEuclidFloor(b, -a, n)
	}

	for i := 0; i < b.N; i++ {
		benchmarkEuclidFloor(b, a, -n)
	}

	for i := 0; i < b.N; i++ {
		benchmarkEuclidFloor(b, -a, -n)
	}
}

func benchmarkEuclidFloor(b *testing.B, a, n int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_, _ = EuclidFloor(a, n)
		}
	}

	return b.Run(fmt.Sprintf("euclidFloor(%d,%d)", a, n), f)
}
