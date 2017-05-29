// https://projecteuler.net/problem=3

package euler

func Problem3() int {
	factor := 2
	composite := 600851475143

	for composite > 2 {
		if (composite % factor) == 0 {
			composite /= factor
		} else {
			factor++
		}
	}

	return factor
}
