package sequence

import (
	"fmt"
	"testing"
)

func TestInts(t *testing.T) {
	its := newInts(
		format{
			newBaseFmt(0, 9),
			newBaseFmt(0, 9),
		},
	)

	for i := 0; i < 10; i++ {
		its.increment()
		fmt.Println(its.current)
	}

	for i := 0; i < 10; i++ {
		its.decrement()
		fmt.Println(its.current)
	}

	t.Fatal()
}
