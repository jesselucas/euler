package main

import (
	"fmt"
	"math"
)

// Special Pythagorean triplet
// https://projecteuler.net/problem=9
func main() {
	// test 0 - 999 to see if the sum of the triplet equals 1000
	for i := 0; i < 1000; i++ {
		triples := pythagoreanTriplets(i)
		for _, t := range triples {
			if t.a+t.b+t.c == 1000 {
				// Print the product of the triplet
				fmt.Println(t.a * t.b * t.c)
				return
			}
		}
	}
}

// triples holds Pythagorean triples
type triple struct {
	a, b, c int
}

// pythagoreanTriplets takes an int and returns it's pythagorean triplet
// https://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples
// r is any int
// s, t = factor-pairs of (r * r)/2
// a = r + s, b = r + t, c = r + s + t
func pythagoreanTriplets(r int) []triple {
	pairs := factorPairs((r * r) / 2)
	triples := []triple{}

	for _, p := range pairs {
		s := p.a
		t := p.b

		a := r + s
		b := r + t
		c := r + s + t

		triples = append(triples, triple{a, b, c})
	}

	return triples
}

// pair holds a factor pair
type pair struct {
	a, b int
}

// generateFactorPairs takes an int and returns its factor pairs
func factorPairs(n int) []pair {
	max := math.Sqrt(float64(n))
	pairs := []pair{}

	for i := 1; i <= int(max); i++ {
		b := n / i    // divide n by a number (i)
		if n%b == 0 { // now see if n is divisible by b
			pairs = append(pairs, pair{i, b})
		}
	}

	return pairs
}
