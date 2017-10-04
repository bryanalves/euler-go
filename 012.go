// https://projecteuler.ent/problem=12

package euler

import "math"

func Problem12() int {
	x := 0
	for x = range triangle() {
		if factorCount(x) > 500 {
			return x
		}
	}

	return x
}

func factorCount(n int) int {
	retval := 0

	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				retval += 1
			} else {
				retval += 2
			}
		}
	}

	return retval
}

func triangle() chan int {
	ch := make(chan int)

	go func() {
		acc := 1
		counter := 2

		for {
			acc += counter
			counter += 1
			ch <- acc
		}

		close(ch)
	}()

	return ch
}
