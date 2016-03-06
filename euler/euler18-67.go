package main

import (
	"./eutil"
	"fmt"
	"io/ioutil"
	"net/http"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func solve(input string) int {
	grid := eutil.ParseGrid(input)
	acc := make([][]int, len(grid))
	for i, _ := range acc {
		acc[i] = make([]int, i+1)
	}
	acc[0][0] = grid[0][0]
	for i, row := range acc {
		for j, element := range row {
			if i+1 >= len(grid) {
				break
			}
			acc[i+1][j+1] = element + grid[i+1][j+1]
			acc[i+1][j] = max(grid[i+1][j]+element, acc[i+1][j])
		}
	}
	maxPath := 0
	for _, elem := range acc[len(grid)-1] {
		if elem > maxPath {
			maxPath = elem
		}
	}
	return maxPath
}

const sample = `
3
7 4
2 4 6
8 5 9 3
`

const round1 = `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`

var round2 = getAll("https://projecteuler.net/project/resources/p067_triangle.txt")

func getAll(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	fmt.Println(solve(sample))
	fmt.Println(solve(round1))
	fmt.Println(solve(round2))
}
