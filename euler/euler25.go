package main

import (
	"fmt"
	"math/big"
)

func solve(digits int) int {
	a := big.NewInt(1)
	b := big.NewInt(1)
	i := 1
	for ; len(a.String()) < digits; i++ {
		c := b
		b = big.NewInt(0).Add(a, b)
		a = c
	}
	return i
}

func main() {
	fmt.Println(solve(1))
	fmt.Println(solve(2))
	fmt.Println(solve(3))
	fmt.Println(solve(1000))
}
