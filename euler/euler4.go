package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(number int) bool {
	return fmt.Sprint(number) == Reverse(fmt.Sprint(number))
}

func PowInt(a int, b int) (res int) {
	res = 1
	for i := 1; i <= b; i++ {
		res *= a
	}
	return
}

func solve(digits int) (res int) {
	from := PowInt(10, digits-1)
	to := PowInt(10, digits)
	for i := to - 1; i >= from; i-- {
		for j := i; j >= from; j-- {
			n := i * j
			if isPalindrome(n) && res < n {
				res = n
			}
		}
	}
	return
}

func main() {
	fmt.Println(solve(3))
}
