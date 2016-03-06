package main

import (
	"fmt"
	"math/big"
)

func solve(power int) int {
	acc := big.NewInt(2)
	two := big.NewInt(2)
	for i := 1; i < power; i++ {
		acc.Mul(acc, two)
	}
	fmt.Println(acc)
	digitAcc := 0
	for _, digit := range acc.String() {
		digitAcc += int(digit) - int('0')
	}
	return digitAcc
}

func main() {
	fmt.Println(solve(15))
	fmt.Println(solve(1000))
}
