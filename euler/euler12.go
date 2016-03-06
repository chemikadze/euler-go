package main

import (
	"./eutil"
	"fmt"
)

func solve(maxDivisors int) int {
	number := 1
	for i := 2; eutil.NumFactors(number) <= maxDivisors; i++ {
		number += i
	}
	return number
}

func main() {
	fmt.Println(solve(5))
	fmt.Println(solve(500))
}
