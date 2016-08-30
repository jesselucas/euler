package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		n      int
		result int
	}{
		{10, 2520},
		{11, 27720},
		{12, 27720},
		{13, 360360},
	}

	for _, test := range tests {
		m := smallestMultiple(test.n)
		if m != test.result {
			t.Fatalf("%v - Received %v. Should be %v", test.n, m, test.result)
		}
	}
}
