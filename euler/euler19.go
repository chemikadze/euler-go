package main

import (
	"fmt"
	"math/big"
)

func solve(power int) int {
	acc := big.NewInt(1)
	for i := 1; i < power; i++ {
		acc.Mul(acc, big.NewInt(int64(i)))
	}
	fmt.Println(acc)
	digitAcc := 0
	for _, digit := range acc.String() {
		digitAcc += int(digit) - int('0')
	}
	return digitAcc
}

func main() {
	fmt.Println(solve(10))
	fmt.Println(solve(100))
}
