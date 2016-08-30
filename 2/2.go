package main

import (
	"fmt"
)

// Even Fibonacci numbers
// https://projecteuler.net/problem=2
func main() {
	v := 0 // value to be added to sum
	sum := 0
	f := fib(1, 2) // start the sequence with 1, 2

	// do not exceed four million
	for v < 4000000 {
		v = f()

		// sum on the even-values
		if v%2 == 0 {
			sum += v
		}
	}

	fmt.Println(sum)

}

// https://en.wikipedia.org/wiki/Fibonacci_number
func fib(a, b int) func() int {
	return func() int {
		r := a // store the start value to be returned

		// every number in the sequence is the sum of the two preceding ones
		a, b = b, a+b // swap so b is the next number in the sequence
		return r
	}
}
