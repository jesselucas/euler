package main

import "testing"

func TestSumPrimes(t *testing.T) {
	var tests = []struct {
		n   int
		sum int
	}{
		{10, 17},
		{11, 17},
		{20, 77},
	}

	for _, test := range tests {
		sum := sumPrimes(test.n)
		if sum != test.sum {
			t.Fatalf("%v - Received %v. Should be %v", test.n, sum, test.sum)
		}
	}
}

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		n      int
		result bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
	}

	for _, test := range tests {
		b := isPrime(test.n)
		if b != test.result {
			t.Fatalf("%v - Received %v. Should be %v", test.n, b, test.result)
		}
	}
}
