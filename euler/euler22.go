package main

import (
	"./eutil"
	"fmt"
	"sort"
	"strings"
)

const url = "https://projecteuler.net/project/resources/p022_names.txt"

func solve(raw string) int {
	names := make([]string, 0, 5000)
	for _, str := range strings.Split(raw, ",") {
		names = append(names, strings.Trim(str, " \n\""))
	}
	sort.Strings(names)
	parsed := make([]int, 0, 5000)
	for _, name := range names {
		acc := 0
		for _, chr := range name {
			acc += int(chr) - int('A') + 1
		}
		parsed = append(parsed, acc)
	}
	acc := 0
	for i, item := range parsed {
		acc += item * (i + 1)
	}
	return acc
}

func main() {
	fmt.Println(solve("\"A\""))
	fmt.Println(solve("\"COLIN\""))
	fmt.Println(solve("\"A\",\"A\""))
	fmt.Println(solve(eutil.GetAll(url)))
}
