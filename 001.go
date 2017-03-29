// https://projecteuler.net/problem=1

package main

func problem_1(n int) int {
    sum := 0
    for i := 0; i < n; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    return sum
}