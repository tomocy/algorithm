package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

func TestSolve(t *testing.T) {
	type input struct {
		vs     []int
		target int
	}
	tests := []struct {
		input    input
		expected bool
	}{
		{
			input{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			true,
		},
		{
			input{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			true,
		},
		{
			input{
				[]int{1, 2, 3, 4, 5},
				5,
			},
			true,
		},
		{
			input{
				[]int{1, 2, 3, 4, 5},
				-3,
			},
			false,
		},
	}

	for _, test := range tests {
		actual := solve(test.input.vs, test.input.target)
		if actual != test.expected {
			algorithm.Reportln(t, "binary search", actual, test.expected)
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve([]int{0, 1, 2, 3, 4, 5}, 4)
	}
}
