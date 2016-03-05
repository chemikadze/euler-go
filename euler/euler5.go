package main

import "fmt"

func solve(max int) int {
	i := max
	for {
		divisible := true
		for j := 1; j <= max; j++ {
			if i%j != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			return i
		}
		i++
	}
}

func main() {
	fmt.Println(solve(20))
}
