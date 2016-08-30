package main

import "testing"

func TestFib(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		result []int
	}{
		{0, 1, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{1, 2, []int{1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, test := range tests {
		f := fib(test.a, test.b)
		v := 0
		i := 0
		for v < test.result[len(test.result)-1] {
			v = f()

			if v != test.result[i] {
				t.Fatalf("Sequence not correct. Received %v. Should be %v", v, test.result[i])
			}

			i++
		}
	}
}
