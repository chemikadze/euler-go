package main

import "fmt"

func solve(max int) int {
	acc := 0
	for i := 1; i <= max; i++ {
		acc += i * i
	}
	sumsq := max * (max + 1) / 2
	return sumsq*sumsq - acc
}

func main() {
	fmt.Println(solve(100))
}
