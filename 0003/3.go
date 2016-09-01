package main

import "fmt"

// Largest prime factor
// https://projecteuler.net/problem=3
func main() {
	var n int64 = 600851475143

	// Create a slice of all primes for n
	p := listPrimes(n)

	// Print out the last prime, which will be the largest
	fmt.Println(p[len(p)-1])
}

// https://en.wikipedia.org/wiki/Prime_number
func listPrimes(n int64) []int64 {
	primes := []int64{} // store primes

	// Starting at 2 (smallest prime)
	for i := int64(2); i < n; i++ {
		// Is i a prime of n?
		if n%i == 0 {
			// Now check it against the other primes to make
			// sure it's not already a factor we have
			for _, p := range primes {
				if i%p == 0 {
					return primes
				}
			}

			// Store the prime
			primes = append(primes, i)
		}
	}

	return nil
}
