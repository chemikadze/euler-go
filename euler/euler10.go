package main

import (
	"fmt"
)

func solve(max int) (res int) {
	primes := make([]bool, max, max)
	primes[0] = true
	primes[1] = true
	for i := 2; i < max; i++ {
		if primes[i] {
			continue
		}
		for j := 2; i*j < max; j++ {
			primes[i*j] = true
		}
	}
	for i, notPrime := range primes {
		if !notPrime {
			res += i
		}
	}
	return
}

func main() {
	fmt.Println(solve(10))
	fmt.Println(solve(2000000))
}
