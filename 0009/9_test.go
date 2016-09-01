package main

import "testing"

func TestPythagoreanTriplets(t *testing.T) {
	var tests = []struct {
		n      int
		result []triple
	}{
		{6, []triple{triple{7, 24, 25}, triple{8, 15, 17}, triple{9, 12, 15}}},
	}

	for _, test := range tests {
		triples := pythagoreanTriplets(test.n)
		for i, triple := range triples {
			if triple.a != test.result[i].a || triple.b != test.result[i].b || triple.c != test.result[i].c {
				t.Fatalf("Pairs not equal %v. Expected %v", triple, test.result[i])
			}
		}
	}
}

func TestFactorPairs(t *testing.T) {
	var tests = []struct {
		n      int
		result []pair
	}{
		{10, []pair{pair{1, 10}, pair{2, 5}}},
		{20, []pair{pair{1, 20}, pair{2, 10}, pair{4, 5}}},
		{18, []pair{pair{1, 18}, pair{2, 9}, pair{3, 6}}},
	}

	for _, test := range tests {
		pairs := factorPairs(test.n)
		for i, p := range pairs {
			if p.a != test.result[i].a || p.b != test.result[i].b {
				t.Fatalf("Pairs not equal %v. Expected %v", p, test.result[i])
			}
		}
	}
}
