package main

import (
	"fmt"
)

// Multiples of 3 and 5
// https://projecteuler.net/problem=1
func main() {
	// Find the sum of all the multiples of 3 or 5 below 1000
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
