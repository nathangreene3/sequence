package main

import "math"

// -------------------------------------------------------------------------
// Definitions and notes
// -------------------------------------------------------------------------
// Addition modulo n: x mod n = r such that x = kn + r for some integer k
// and r in [0,n). r = x % n and k = [x/n].
// -------------------------------------------------------------------------
// Zn: Let n be a non-zero integer. If n > 0, let Zn denote the set
// {0, 1, ..., n-1} under addition modulo n. Otherwise, if n < 0, let Zn
// denote the set {n+1, ..., -1, 0} under addition modulo n.
// -------------------------------------------------------------------------
// The values "carry" and "borrow" refer to k and are useful in addition
// and subtraction over the external direct product of several sets of Zn.
// See J.A. Gallian's Contemporary Abstract Algebra, 6th Ed., chapters 1, 2,
// and 8.
// -------------------------------------------------------------------------
// A possible property: if x >= 0 and n > 0, then x mod -n = -(-x mod n).
// -------------------------------------------------------------------------
// Links and sources
// * Contemporary Abstract Algebra, 6th Ed., by J.A. Gallian
// * https://en.wikipedia.org/wiki/Euclidean_division
// * https://en.wikipedia.org/wiki/Modular_arithmetic
// * https://en.wikipedia.org/wiki/Modulo_operation
// -------------------------------------------------------------------------

// addMod returns r = (a+b) mod n with the carried amount k such that a + b = kn + r.
//
// * r: The sum modulus n. The quantity a+b is expressed as a multiple of the modulus plus some residual r. Here, r is chosen to share the sign of n such that |r| < |n|. This is equivalent to
// * sgn(r) = sgn(n)
// * sgn(k) = sgn(a+b)
func addMod(a, b, modulus int) (int, int) {
	var (
		ka, ra = EuclidFloor(a, modulus)
		kb, rb = EuclidFloor(b, modulus)
		k, r   = EuclidFloor(ra+rb, modulus)
	)

	return r, ka + kb + k
}

// subtractMod returns r = (a-b) mod n with the borrowed amount k
// such that a - b = kn - r for |r| < |n|.
func subtractMod(a, b, modulus int) (int, int) {
	var (
		ka, ra = EuclidFloor(a, modulus)
		kb, rb = EuclidFloor(b, modulus)
		k, r   = EuclidFloor(ra-rb, modulus)
	)

	return r, -ka + kb - k
}

// EuclidFloor returns (k,r) such that x = kn + r for a given non-zero
// modulus n.
//
// * |r| < |n|
// * sgn(r) = sgn(n)
// * sgn(k) = sgn(x)
func EuclidFloor(x, modulus int) (int, int) {
	r := (x%modulus + modulus) % modulus // k := int(math.Floor(float64(x) / float64(modulus)))
	return (x - r) / modulus, r          // return k, x - k*modulus
}

// EuclidFloor2 ...
func EuclidFloor2(x, modulus int) (int, int) {
	if x == 0 {
		return 0, 0
	}

	r := (x%modulus + modulus) % modulus // k := int(math.Floor(float64(x) / float64(modulus)))
	return (x - r) / modulus, r          // return k, x - k*modulus
}

// EuclidTrunc ...
func EuclidTrunc(x, modulus int) (int, int) {
	k := x / modulus
	return k, x - k*modulus
}

// Euclid satisfies the original Euclidean division algorithm definition in which (k,r) are returned such that x = k*n + r and 0 <= r < |n|.
func Euclid(x, modulus int) (int, int) {
	if modulus < 0 {
		modulus *= -1
	}

	k := int(math.Floor(float64(x) / float64(modulus)))
	return k, x - k*modulus
}

// Euclid2 ...
func Euclid2(x, modulus int) (int, int) {
	if modulus < 0 {
		modulus *= -1
	}

	var k int
	f := float64(x) / float64(modulus)
	if 0 < f {
		k = int(f)
	} else if f == 0 {
		return 0, 0
	} else {

	}

	return k, x - k*modulus
}
