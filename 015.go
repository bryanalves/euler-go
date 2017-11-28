// https://projecteuler.net/problem=15

package euler

import "math/big"

func Problem15() int64 {
	grid := int64(20)
	retval := big.NewInt(0).Binomial(grid*2, grid)
	return retval.Int64()
}
