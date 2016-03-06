package main

import "fmt"

func isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

func solve(sum int) int {
	for a := 1; a < sum; a++ {
		for b := 1; b < a; b++ {
			for c := 1; c < b; c++ {
				if isTriplet(c, b, a) && a+b+c == sum {
					return a * b * c
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(solve(12))
	fmt.Println(solve(1000))
}
