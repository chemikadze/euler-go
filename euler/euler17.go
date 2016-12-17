package main

import (
	"fmt"
	"unicode/utf8"
)

var simpleNumbers = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var teenNumbers = []string{
	"",
	"",
	"twen",
	"thir",
	"four",
	"fif",
	"six",
	"seven",
	"eigh",
	"nine",
}

var tyNumbers = []string{
	"",
	"",
	"twen",
	"thir",
	"for",
	"fif",
	"six",
	"seven",
	"eigh",
	"nine",
}

func numberToHuman(number int) []string {
	acc := make([]string, 0)
	if number/1000 != 0 {
		acc = append(acc, simpleNumbers[number/1000])
		acc = append(acc, "thousand")
	}
	if number/100%10 != 0 {
		acc = append(acc, simpleNumbers[number/100%10])
		acc = append(acc, "hundred")
	}
	lowerTwo := number % 100
	firstDigit := lowerTwo / 10
	secondDigit := lowerTwo % 10
	if lowerTwo != 0 {
		if len(acc) > 0 {
			acc = append(acc, "and")
		}
		if lowerTwo <= 12 {
			acc = append(acc, simpleNumbers[lowerTwo])
		} else if lowerTwo > 12 && lowerTwo < 20 {
			acc = append(acc, teenNumbers[secondDigit]+"teen")
		} else {
			acc = append(acc, tyNumbers[firstDigit]+"ty")
			if secondDigit != 0 {
				acc = append(acc, simpleNumbers[secondDigit])
			}
		}
	}
	return acc
}

func countLetters(text []string) int {
	letters := 0
	for _, word := range text {
		letters += utf8.RuneCountInString(word)
	}
	return letters
}

func solve(number int) int {
	acc := 0
	for i := 1; i <= number; i++ {
		acc += countLetters(numberToHuman(i))
	}
	return acc
}

func main() {
	fmt.Println(numberToHuman(1))
	fmt.Println(numberToHuman(2))
	fmt.Println(numberToHuman(10))
	fmt.Println(numberToHuman(25))
	fmt.Println(numberToHuman(125))
	fmt.Println(numberToHuman(7125))
	fmt.Println(solve(5))
	fmt.Println(solve(1000))
}
