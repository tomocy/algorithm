package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{5, 120},
		{1, 1},
		{0, 1},
	}

	for _, test := range tests {
		actual := solve(test.input)
		if actual != test.expected {
			t.Errorf("unexpected result: got %d, expect %d\n", actual, test.expected)
		}
	}
}
