// https://projecteuler.net/problem=10

package euler

func Problem10() int {
	ch := make(chan int)
	go prime_sieve(ch, 2000000)

	ret := 0
	i := 0
	ok := true
	for {
		i, ok = <-ch
		if !ok {
			break
		}
		ret += i
	}

	return ret
}

func prime_sieve(ch chan int, max int) {
	b := make([]bool, max)
	for i := 2; i < max; i++ {
		if b[i] == true {
			continue
		}
		ch <- i
		for k := i * i; k < max; k += i {
			b[k] = true
		}
	}
	close(ch)
}
