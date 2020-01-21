package zmodn

import (
	"fmt"
	"testing"

	"github.com/nathangreene3/math"
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
			answer, borrow := subtractWithBorrow(test.x, test.y, test.n)
			if test.answer != answer || test.borrow != borrow {
				t.Fatalf("\n(%d - %d) mod %d\nexpected (%d,%d)\nreceived (%d,%d)\n", test.x, test.y, test.n, test.answer, test.borrow, answer, borrow)
			}
		} else {
			answer, carry := addWithCarry(test.x, test.y, test.n)
			if test.answer != answer || test.carry != carry {
				t.Fatalf("\n(%d + %d) mod %d\nexpected (%d,%d)\nreceived (%d,%d)\n", test.x, test.y, test.n, test.answer, test.carry, answer, carry)
			}
		}
	}
}

func TestEuclidsCoeffs(t *testing.T) {
	tests := []struct {
		x, n, k, r int
	}{
		// 0 < n
		{
			// -6 =-2*3 + 0
			x: -6,
			n: 3,
			k: -2,
			r: 0,
		},
		{
			// -5 = -2*3 + 1
			x: -5,
			n: 3,
			k: -2,
			r: 1,
		},
		{
			// -4 = -2*3 + 2
			x: -4,
			n: 3,
			k: -2,
			r: 2,
		},
		{
			// -3 = -1*3 + 0
			x: -3,
			n: 3,
			k: -1,
			r: 0,
		},
		{
			// -2 = -1*3 + 1
			x: -2,
			n: 3,
			k: -1,
			r: 1,
		},
		{
			// -1 = -1*3 + 2
			x: -1,
			n: 3,
			k: -1,
			r: 2,
		},
		{
			// 0 = 0*3 + 0
			x: 0,
			n: 3,
			k: 0,
			r: 0,
		},
		{
			// 1 = 0*3 + 1
			x: 1,
			n: 3,
			k: 0,
			r: 1,
		},
		{
			// 2 = 0*3 + 2
			x: 2,
			n: 3,
			k: 0,
			r: 2,
		},
		{
			// 3 = 1*3 + 0
			x: 3,
			n: 3,
			k: 1,
			r: 0,
		},
		{
			// 4 = 1*3 + 1
			x: 4,
			n: 3,
			k: 1,
			r: 1,
		},
		{
			// 5 = 1*3 + 2
			x: 5,
			n: 3,
			k: 1,
			r: 2,
		},
		{
			// 6 = 2*3 + 0
			x: 6,
			n: 3,
			k: 2,
			r: 0,
		},

		// n < 0
		{
			// -6 = 2*-3 - 0
			x: -6,
			n: -3,
			k: 2,
			r: 0,
		},
		{
			// -5 = 1*-3 - 2
			x: -5,
			n: -3,
			k: 1,
			r: -2,
		},
		{
			x: -4,
			n: -3,
			k: 1,
			r: -1,
		},
		{
			x: -3,
			n: -3,
			k: 1,
			r: 0,
		},
		{
			x: -2,
			n: -3,
			k: 0,
			r: -2,
		},
		{
			x: -1,
			n: -3,
			k: 0,
			r: -1,
		},
		{
			x: 0,
			n: -3,
			k: 0,
			r: 0,
		},
		{
			x: 1,
			n: -3,
			k: -1,
			r: -2,
		},
		{
			x: 2,
			n: -3,
			k: -1,
			r: -1,
		},
		{
			x: 3,
			n: -3,
			k: -1,
			r: 0,
		},
		{
			x: 4,
			n: -3,
			k: -2,
			r: -2,
		},
		{
			x: 5,
			n: -3,
			k: -2,
			r: -1,
		},
		{
			x: 6,
			n: -3,
			k: -2,
			r: 0,
		},
	}

	for _, test := range tests {
		k, r := euclidsCoeffs(test.x, test.n)
		if test.k != k || test.r != r {
			t.Fatalf("\ngiven (%d,%d)\nexpected (%d,%d)\nreceived (%d,%d)\n", test.x, test.n, test.k, test.r, k, r)
		}
	}
}

func TestBase(t *testing.T) {
	fmt.Println(math.Base(16, 3))
	fmt.Println(math.Base(8, 3))
	fmt.Println(math.Base(24, 3))
	t.Fatal()
}
