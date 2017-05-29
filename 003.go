// https://projecteuler.net/problem=3

package euler

func Problem3(composite int) int {
	factor := 2

	for composite > 2 {
		if (composite % factor) == 0 {
			composite /= factor
		} else {
			factor++
		}
	}

	return factor
}
