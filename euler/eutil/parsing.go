package eutil

import (
	"strconv"
	"strings"
)

func ParseGrid(raw string) [][]int {
	result := make([][]int, 0)
	for _, line := range strings.Split(raw, "\n") {
		trimmedLine := strings.Trim(line, " ")
		if trimmedLine == "" {
			continue
		}
		row := make([]int, 0)
		for _, item := range strings.Split(trimmedLine, " ") {
			num, _ := strconv.Atoi(item)
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}
