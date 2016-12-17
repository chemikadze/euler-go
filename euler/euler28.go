package main

import (
	"fmt"
)

func solve(size int) int {
	// prepare
	matrix := make([][]int, size+1)
	for i := range matrix {
		matrix[i] = make([]int, size+1)
	}
	checkVectors := [][]int{
		[]int{+0, -1},
		[]int{-1, +0},
		[]int{+0, +1},
		[]int{+1, +0},
	}
	moveVectors := [][]int{
		[]int{+1, +0},
		[]int{+0, -1},
		[]int{-1, +0},
		[]int{+0, +1},
	}
	// fill
	i, j := size/2, size/2
	if size == 1 {
		i, j = 0, 0
	}
	counter := 1
	matrix[i][j] = counter
	counter += 1
	direction := 3
	for i < size && j < size {
		checkI := checkVectors[direction][0]
		checkJ := checkVectors[direction][1]
		moveI := moveVectors[direction][0]
		moveJ := moveVectors[direction][1]
		for {
			i += moveI
			j += moveJ
			matrix[i][j] = counter
			counter += 1
			if matrix[i+checkI][j+checkJ] == 0 {
				break
			}
		}
		direction = (direction + 1) % 4
	}
	// count
	sum := 0
	for i := 0; i < size; i++ {
		sum += matrix[i][i] + matrix[i][size-i-1]
	}
	return sum - 1
}

func main() {
	fmt.Println(solve(1))
	fmt.Println(solve(3))
	fmt.Println(solve(5))
	fmt.Println(solve(1001))
	return
}
