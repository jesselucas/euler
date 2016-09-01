package main

import (
	"fmt"
	"math"
)

// Summation of primes
// https://projecteuler.net/problem=10
func main() {
	fmt.Println(sumPrimes(2000000))
}

// sumPrimes sum of the primes below n
func sumPrimes(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	return sum
}

// https://en.wikipedia.org/wiki/Prime_number
func isPrime(n int) bool {
	// Anything smaller 2 isn't prime
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		// no positive divisors other than 1 and itself
		if n%i == 0 {
			return false
		}
	}

	return true
}
