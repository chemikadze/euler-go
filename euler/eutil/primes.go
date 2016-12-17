package eutil

// import (
// 	"fmt"
// )

var primeNumbers = make([]int, 0)

var maxprime = 1000000 //math.MaxInt32
var primes = make([]bool, maxprime)

func init() {
	primes[0] = true
	primes[1] = true
	for i := 2; i < maxprime; i++ {
		if primes[i] {
			continue
		}
		for j := 2; i*j < maxprime; j++ {
			primes[i*j] = true
		}
	}
	for i, notPrime := range primes {
		if !notPrime {
			primeNumbers = append(primeNumbers, i)
		}
		primes[i] = !primes[i]
	}
	return
}

type Factor struct {
	Prime int
	Count int
}

func IsPrime(number int) bool {
	if number < 0 {
		number = -number
	}
	if number < maxprime {
		return primes[number]
	} else {
		for i := 2; i < number/2; i += 1 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

func PrimeFactors(number int) []Factor {
	factors := make([]Factor, 0, 10)
	for _, prime := range primeNumbers {
		if prime > number {
			break
		}
		if number%prime == 0 {
			count := 0
			num := number
			for num%prime == 0 {
				count++
				num /= prime
			}
			factors = append(factors, Factor{prime, count})
		}
	}
	return factors
}

func NumFactors(number int) int {
	factors := PrimeFactors(number)
	p := 1
	for _, factor := range factors {
		p *= (factor.Count + 1)
	}
	return p
}
