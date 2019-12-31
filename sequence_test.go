package sequence

import (
	"fmt"
	"testing"
)

func TestInts(t *testing.T) {
	its := newInts(
		format{
			newBaseFmt(0, 3),
			newBaseFmt(0, 3),
		},
	)

	for i := 0; i < 4; i++ {
		its.increment()
		fmt.Println(its.current)
	}

	for i := 0; i < 4; i++ {
		its.decrement()
		fmt.Println(its.current)
	}

	t.Fatal()
}

func TestAddSubtract(t *testing.T) {
	its := newInts(newFormat(newBaseFmt(1, 3), newBaseFmt(2, 3)))
	fmt.Printf("%v\n", its.current)

	for i := 0; i < 5; i++ {
		its.increment()
		fmt.Printf("%v\n", its.current)
	}

	fmt.Println()
	for i := 0; i < 5; i++ {
		its.decrement()
		fmt.Printf("%v\n", its.current)
	}

	t.Fatal()
}
