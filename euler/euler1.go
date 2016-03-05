package main

import "fmt"

func solve(max int) (acc int) {
	for i := 0; i < max; i++ {
		if i%5 == 0 || i%3 == 0 {
			acc += i
		}
	}
	return
}

func main() {
	fmt.Println(solve(1000))
}
