// https://projecteuler.net/problem=14

package euler

func Problem14() int {
	longest_chain := 0
	longest_number := 0

	for i := 1; i < 1000000; i++ {
		length := collatzLength(i)
		if length > longest_chain {
			longest_chain = length
			longest_number = i
		}
	}

	return longest_number
}

func collatzLength(n int) int {
	retval := 0

	for ; n > 1; retval++ {
		if (n % 2) == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}

	return retval
}
