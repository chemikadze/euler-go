package main

import "fmt"

func solve() (acc int) {
	a, b := 1, 1
	for a < 4000000 {
		if a%2 == 0 {
			acc += a
		}
		a, b = b, a+b
	}
	return
}

func main() {
	fmt.Println(solve())
}
