package main

import "testing"

func TestListPrimes(t *testing.T) {
	var tests = []struct {
		n      int64
		result []int64
	}{
		{13195, []int64{5, 7, 13, 29}},
		{600851475143, []int64{71, 839, 1471, 6857}},
	}

	for _, test := range tests {
		primes := listPrimes(test.n)
		for i, p := range primes {
			if p != test.result[i] {
				t.Fatalf("Sequence not correct. Received %v. Should be %v", p, test.result[i])
			}
		}
	}
}
