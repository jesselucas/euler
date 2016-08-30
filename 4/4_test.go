package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		n      int
		result bool
	}{
		{9009, true},
		{919, true},
		{101, true},
		{100, false},
		{10, false},
		{11, true},
	}

	for _, test := range tests {
		p := isPalindrome(test.n)
		if p != test.result {
			t.Fatalf("%v - Received %v. Should be %v", test.n, p, test.result)
		}
	}
}
