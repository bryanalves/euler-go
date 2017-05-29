// https://projecteuler.net/problem=5

package main

func problem_5() int {
	ret := lcm(1, 2)

	for i := 3; i <= 20; i++ {
		ret = lcm(ret, i)
	}

	return ret
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
