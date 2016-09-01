package main

import "fmt"

// Smallest multiple
// https://projecteuler.net/problem=5
func main() {
	fmt.Println(smallestMultiple(20))
}

// smallestMultiple keeps looping until 1 - n is divisible into a number
func smallestMultiple(n int) int {
	success := 0 // we count how many successes
	result := 1  // start the result to 1

	for {
		// Test 1- n
		for i := 1; i <= n; i++ {
			// If it is divisible then increment success
			if result%i == 0 {
				success++
			} else {
				// Restart and break if not successful
				success = 0
				break
			}
		}

		// If we finally have a winner then return
		if success >= n {
			return result
		}

		result++
	}
}
