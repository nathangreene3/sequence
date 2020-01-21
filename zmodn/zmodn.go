package zmodn

// Zn = {0, 1, ..., n-1},  n > 0
//    = {n+1, ..., -1, 0}, n < 0 (my extension to the definition)

// Addition modulo n: x mod n = r such that x = kn + r for some
// integer k and r in [0,n). r = x % n and k = [x/n].

// The values "carry" and "borrow" refer to k and are useful in
// addition and subtraction over the external direct product of
// several sets of Zn. See J.A. Gallian's Contemporary Abstract
// Algebra, 6th Ed., chapters 1, 2, and 8.

// A possible property: if x >= 0 and n > 0, then x mod -n = -(-x mod n).

// addWithCarry returns r = (a+b) mod n with the carried amount k such that a+b = kn+r.
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

	return r, -ka + kb - k
}

// euclidsCoeffs returns (k,r) such that x = kn + r for a given
// modulus n != 0. If n > 0, then 0 <= r < n. Otherwise, n < r <= 0.
// In either case, k = (x-r)/n.
func euclidsCoeffs(x, modulus int) (k int, r int) {
	if modulus == 0 {
		panic("")
	}

	r = (x%modulus + modulus) % modulus
	return (x - r) / modulus, r
}
