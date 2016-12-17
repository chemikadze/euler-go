package main

import (
	"./eutil"
	"fmt"
)

func seqLength(a int, b int) int {
	number := 0
	for ; eutil.IsPrime(number*number + number*a + b); number++ {
	}
	return number
}

func solve() int {
	maxMult := 41
	maxLength := 40
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			length := seqLength(a, b)
			if length > maxLength {
				maxLength = length
				maxMult = a * b
			}
		}
	}
	return maxMult
}

func main() {
	fmt.Println(seqLength(1, 41))
	fmt.Println(seqLength(-79, 1601))
	fmt.Println(solve())
	return
}
