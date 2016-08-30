package main

import (
	"fmt"
	"strconv"
)

// Largest palindrome product
// https://projecteuler.net/problem=4
func main() {
	result := 0

	// For 100 - 999
	for i := 100; i < 999; i++ {
		// Now for each 3 digit create another number 100 - 999
		for j := 100; j < 999; j++ {
			t := i * j
			if isPalindrome(t) {
				// if the new product was bigger than the last store it
				if t > result {
					result = t
				}
			}
		}
	}

	// The biggest result!
	fmt.Println(result)
}

// https://en.wikipedia.org/wiki/Palindromic_number
func isPalindrome(n int) bool {
	// Convert to string then to byte slice
	bs := []byte(strconv.Itoa(n))
	l := len(bs)
	for i, b := range bs {
		// Test the beginning of the slice to the end
		if b != bs[l-(i+1)] {
			return false
		}
	}

	return true
}
