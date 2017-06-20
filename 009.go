// https://projecteuler.net/problem=9

package euler

import "math"

func Problem9() int {
	c := 0
	for a := 1; a < 998; a++ {
		for b := 1; b < 999-a; b++ {
			c = int(math.Sqrt(float64(a*a + b*b)))
			if (a+b+c == 1000) && pythagorean_match(a, b, c) {
				return a * b * c
			}
		}
	}

	return 0
}

func pythagorean_match(a int, b int, c int) bool {
	return ((a*a)+(b*b) == (c * c))
}
