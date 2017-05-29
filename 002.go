// https://projecteuler.net/problem=2

package main

func problem_2() int {
	fib_gen := fibonacci()
	acc, fib_current := 0, 0

	for fib_current < 4000000 {
		fib_current = fib_gen()
		if (fib_current % 2) == 0 {
			acc += fib_current
		}
	}

	return acc
}

func fibonacci() func() int {
	x, y := 1, 0
	return func() int {
		z := x + y
		x = y
		y = z
		return x
	}
}
