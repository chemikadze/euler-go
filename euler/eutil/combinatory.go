package eutil

func Binomial(n, k int) int {
	if k > n {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return Binomial(n-1, k-1) + Binomial(n-1, k)
}
