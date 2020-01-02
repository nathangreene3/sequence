package sequence

// baseFmt ...
type baseFmt struct {
	min, max int
}

// newBaseFmt ...
func newBaseFmt(min, max int) baseFmt {
	if max < min || min < 0 {
		panic("minimum must be non-negative and less than maximum")
	}

	return baseFmt{min: min, max: max}
}

// addWithCarry returns (a+b)mod(max+1) + min.
func (bf *baseFmt) addWithCarry(a, b int) (int, int) {
	var (
		c     = a + b
		carry int
	)

	for ; bf.max < c; carry++ {
		c -= bf.max - bf.min + 1
	}

	return c, carry
}

func (bf *baseFmt) subtractWithBorrow(a, b int) (int, int) {
	if a < b {
		return b - a + bf.min, 0
	}

	return bf.addWithCarry(a, bf.max-bf.min-b+1)
}
