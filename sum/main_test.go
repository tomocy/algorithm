package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 15},
	}

	for _, test := range tests {
		actual := solve(test.input)
		if actual != test.expected {
			algorithm.Reportln(t, "sum", actual, test.expected)
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve([]int{0, 1, 2, 3, 4, 5})
	}
}
