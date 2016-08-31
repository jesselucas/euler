package main

import "fmt"

// 10001st prime
// https://projecteuler.net/problem=7
func main() {
	count := 0
	i := 2
	primeNum := 10001 // Index of prime number
	for {
		if isPrime(i) {
			count++ // if we found a prime increment count
		}

		// if the count reaches the prime number index we're looking for break out of the loop
		if count >= primeNum {
			break
		}

		i++
	}

	fmt.Println(i)
}

// https://en.wikipedia.org/wiki/Prime_number
func isPrime(n int) bool {
	// Anything smaller 2 isn't prime
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		// no positive divisors other than 1 and itself
		if n%i == 0 {
			return false
		}
	}

	return true
}
