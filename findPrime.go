package main

import "math"

func isPrime(input int) bool {
	maxFactor := int(math.Sqrt(float64(input)))
	for i := 2; i <= maxFactor; i++ {
		if input%i == 0 {
			return false
		}
	}

	return true
}

func FindPrime(input int) int {
	var r int
	for i := input; i >= 2; i-- {
		if isPrime(i) {
			r = i
			break
		}
	}

	return r
}
