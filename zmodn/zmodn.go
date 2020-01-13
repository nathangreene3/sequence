package zmodn

type zInts []zInt

func newZInts(zs ...zInt) zInts {
	z := make(zInts, len(zs))
	copy(z, zs)
	return z
}

// zInt is an element of Zn = {0,1,...,n-1}.
type zInt struct {
	value, modulus int
}

func newZInt(value, modulus int) zInt {
	if value < 0 || modulus < 1 {
		panic("argument out of range")
	}

	return zInt{value: value, modulus: modulus}
}

func (z *zInt) addWithCarry(x int) int {
	v, carry := addWithCarry(z.value, x, z.modulus)
	z.value = v
	return carry
}

func (z *zInt) subtractWithBorrow(x int) int {
	v, borrow := subtractWithBorrow(z.value, x, z.modulus)
	z.value = v
	return borrow
}

// Zn = {0,1,...,n-1}

// Addition modulo n: x mod n = r such that x = qn + r for some
// integer q and r in [0,n). r = x % n and q = [x/n].

// The values "carry" and "borrow" refer to q and are useful in
// addition and subtraction over the external direct product of
// several sets of Zn. That is, to compute 11-02 over Z3^2, 1-2 is
// computed over Z3 borrowing 1 from the left-most 1, then 0-0 is
// computed as simply 0 resulting in 11-02 --> 04-02 = 02. If that
// isn't clear, then see J.A. Gallian's Contemporary Abstract
// Algebra, 6th Ed., chapters 1, 2, and 8.

// addWithCarry returns (a+b) mod (modulus) with the carried amount.
func addWithCarry(a, b, modulus int) (int, int) {
	var (
		ka, ra = euclidsCoeffs(a, modulus)
		kb, rb = euclidsCoeffs(b, modulus)
		k, r   = euclidsCoeffs(ra+rb, modulus)
	)

	return r, ka + kb + k
}

// subtractWithBorrow ...
func subtractWithBorrow(a, b, modulus int) (int, int) {
	var (
		ka, ra = euclidsCoeffs(a, modulus)
		kb, rb = euclidsCoeffs(b, modulus)
		k, r   = euclidsCoeffs(ra-rb, modulus)
	)

	return r, -k - ka - kb
}

func euclidsCoeffs(x, modulus int) (k int, r int) {
	if modulus == 0 {
		panic("")
	}

	r = (x%modulus + modulus) % modulus
	return (x - r) / modulus, r
}
