// https://projecteuler.net/problem=7

package euler

import "math"

func Problem7() int {
	x := 0
	ch := make(chan int)
	go prime_generate(ch)

	for i := 0; i < 10001; i++ {
		x = <-ch
	}

	return x
}

func prime_generate(ch chan int) {
	ch <- 2

	for i := 3; ; i += 2 {
		if is_prime(i) {
			ch <- i
		}
	}
}

func is_prime(n int) bool {
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
