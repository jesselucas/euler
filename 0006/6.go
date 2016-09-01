package main

import "fmt"

// Sum square difference
// https://projecteuler.net/problem=6
func main() {
	fmt.Println(diffSumSquares(100))
}

func diffSumSquares(n int) int {
	var sumSquares int // the sum of the squares
	var squareSums int // the square of the sum
	for i := 1; i <= n; i++ {
		sq := i * i
		sumSquares += sq
		squareSums += i
	}
	squareSums *= squareSums

	return squareSums - sumSquares
}
