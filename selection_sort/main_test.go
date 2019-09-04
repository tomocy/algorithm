package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected []int
	}{
		{[]int{5, 3, 6, 2, 1}, []int{1, 2, 3, 5, 6}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		actual := solve(test.input)
		if len(actual) != len(test.expected) {
			t.Fatalf("unexpected length: got %d, expect %d\n", len(actual), len(test.expected))
		}
		for i, expected := range test.expected {
			if actual[i] != expected {
				algorithm.Reportln(t, "selection sort", actual[i], expected)
			}
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve([]int{3, 5, 2, 4, 1})
	}
}
