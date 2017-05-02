// https://projecteuler.net/problem=6

package main

func problem_6() int {
    return square_sums(100) - sum_squares(100)
}

func sum_squares(n int) int {
    ret := 0
    for i := 1 ; i <= n ; i++ {
        ret += (i * i)
    }

    return ret
}

func square_sums(n int) int {
    ret := 0
    for i := 1 ; i <= n ; i++ {
        ret += i
    }

    return ret * ret
}
