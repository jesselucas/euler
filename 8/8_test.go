package main

import "testing"

func TestLargestProduct(t *testing.T) {
	var tests = []struct {
		n      int
		result int
	}{
		{2, 81},
		{3, 648},
		{4, 5832},
		{5, 40824},
	}

	for _, test := range tests {
		r := largestProduct(digits, test.n)
		if r != test.result {
			t.Fatalf("%v - Received %v. Should be %v", test.n, r, test.result)
		}
	}
}
