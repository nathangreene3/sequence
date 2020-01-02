package zmodn

import (
	"fmt"
	"testing"
)

func TestAddWithCarry(t *testing.T) {
	n := 3
	for x := 0; x < 2*n; x++ {
		for y := 0; y < 2*n; y++ {
			z, c := addWithCarry(x, y, n)
			fmt.Printf("%d + %d = %d --> %d, carry: %d\n", x, y, x+y, z, c)
		}
	}

	for x := 0; x < 2*n; x++ {
		for y := 0; y < 2*n; y++ {
			z, c := subtractWithBorrow(x, y, n)
			fmt.Printf("%d - %d = %d --> %d, borrow: %d\n", x, y, x-y, z, c)
		}
	}

	x, y, n := 1, -2, 3
	z, c := addWithCarry(x, y, n)
	fmt.Printf("%d + %d = %d --> %d,  carry: %d\n", x, y, x+y, z, c)
	z, c = subtractWithBorrow(x, y, n)
	fmt.Printf("%d - %d = %d --> %d, borrow: %d\n", x, y, x-y, z, c)

	y *= -1
	z, c = addWithCarry(x, y, n)
	fmt.Printf("%d + %d = %d --> %d,  carry: %d\n", x, y, x+y, z, c)
	z, c = subtractWithBorrow(x, y, n)
	fmt.Printf("%d - %d = %d --> %d, borrow: %d\n", x, y, x-y, z, c)

	t.Fatal()
}

func TestZIntAddWithCarry(t *testing.T) {
	n := 3
	z := newZInt(0, n)

	fmt.Printf("z = %d\n", z.value)
	for x := 0; x < 2*n; x++ {
		oldZ := z.value
		c := z.addWithCarry(x)
		fmt.Printf("%d + %d --> %d,  carry: %d\n", oldZ, x, z.value, c)
	}

	for x := 0; x < 2*n; x++ {
		oldZ := z.value
		b := z.subtractWithBorrow(x)
		fmt.Printf("%d - %d --> %d, borrow: %d\n", oldZ, x, z.value, b)
	}

	t.Fatal()
}
