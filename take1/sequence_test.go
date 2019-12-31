package sequence

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var (
		f = NewFormat(
			NewCharFmt(1, 8, 2),
			NewCharFmt(0, 9, 1),
			NewCharFmt(1, 2, 0),
		)
	)

	fmt.Printf("%v", f)
	t.Fatal()
}
