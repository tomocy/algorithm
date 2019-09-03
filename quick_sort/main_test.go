package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 4, 1, 5, -8}, []int{-8, 1, 3, 4, 5}},
		{},
	}

	for _, test := range tests {
		actual := solve(test.input)
		if len(actual) != len(test.expected) {
			t.Fatalf("unexpected length: got %d, expect %d\n", len(actual), len(test.expected))
		}
		for i, expected := range test.expected {
			if actual[i] != expected {
				algorithm.Reportln(t, "quick sort", actual[i], expected)
			}
		}
	}
}
