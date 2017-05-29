// https://projecteuler.net/problem=4

package euler

import "strconv"

func Problem4() int {
	max := 1

	var product int

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			product = i * j
			if (product > max) && is_palindrome(product) {
				max = product
			}
		}
	}
	return max
}

func is_palindrome(n int) bool {
	intstring := strconv.Itoa(n)
	return Reverse(intstring) == intstring
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
