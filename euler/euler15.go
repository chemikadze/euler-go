package main

import (
	"fmt"
)

func solve(cells int) int {
	field := make([][]int, cells+1)
	for i := range field {
		field[i] = make([]int, cells+1)
	}
	field[0][0] = 1
	for i, row := range field {
		for j := range row {
			if i-1 >= 0 {
				field[i][j] += field[i-1][j]
			}
			if j-1 >= 0 {
				field[i][j] += field[i][j-1]
			}
		}
	}
	return field[cells][cells]
}

func main() {
	fmt.Println(solve(2))
	fmt.Println(solve(20))
}
