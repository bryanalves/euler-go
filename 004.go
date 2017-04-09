// https://projecteuler.net/problem=4

package main

import "strconv"
//import "fmt"

func problem_4() int {
    max := 1
    t := 0

    for i := 0; i <= 999; i++ {
        for j:= 0; j <= 999; j++ {
            t = i * j
            if (t > max) && is_palindrome(t) {
                max = t
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
