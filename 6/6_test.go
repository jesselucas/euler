package main

import "testing"

func TestDiffSumSquares(t *testing.T) {
	var tests = []struct {
		n      int
		result int
	}{
		{10, 2640},
		{11, 3850},
		{12, 5434},
		{13, 7462},
	}

	for _, test := range tests {
		s := diffSumSquares(test.n)
		if s != test.result {
			t.Fatalf("%v - Received %v. Should be %v", test.n, s, test.result)
		}
	}
}
