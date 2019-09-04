package main

import (
	"fmt"
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func([]int) int{
	"solve": solve, "solve2": solve2,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 15},
	}

	for _, test := range tests {
		for name, s := range solutions {
			actual := s(test.input)
			if actual != test.expected {
				algorithm.Reportln(t, fmt.Sprintf("sum of %s", name), actual, test.expected)
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{0, 1, 2, 3, 4, 5})
			}
		})
	}
}
