package main

import "fmt"

func solve(number int64) (res int64) {
	res = 1
	var i int64 = 2
	for ; i <= number; i++ {
		for number%i == 0 {
			number /= i
			res = i
		}
	}
	return
}

func main() {
	fmt.Println(solve(600851475143))
}
