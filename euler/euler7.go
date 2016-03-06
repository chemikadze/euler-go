package main

import (
	"fmt"
)

func isPrime(x int) bool {
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func solve(max int) int {
	current := 2
	primeIndex := 0
	for {
		if isPrime(current) {
			primeIndex += 1
		}
		if primeIndex == max {
			break
		}
		current += 1
	}
	return current
}

func main() {
	fmt.Println(solve(1))
	fmt.Println(solve(2))
	fmt.Println(solve(3))
	fmt.Println(solve(6))
	fmt.Println(solve(10001))
}
