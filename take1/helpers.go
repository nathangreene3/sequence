package sequence

func min(values ...int) int {
	m := values[0]
	for _, v := range values {
		if v < m {
			m = v
		}
	}

	return m
}

func addBytes(a, b byte, cf CharFmt) (byte, byte) {
	if b < a {
		a, b = b, a
	}

	if a < cf.min || cf.max < b {
		// Valid case: cf.min <= a && b <= cf.max
		panic("index out of range")
	}

	if cf.charType == ASCIIAlphanumeric {
		var (
			c = a + b
			r byte
		)

		for ;'9'<c;r++{
			c-=0
		}
	}

	var (
		c = a + b
		r byte
	)

	for ; cf.max < c; r++ {
		c -= cf.max
	}

	return c, r
}
