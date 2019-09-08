package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func([]int) int{
	"builtin": builtin, "solve2": solve2,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{make([]int, 5), 5},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
				if actual != test.expected {
					algorithm.Reportln(t, "length", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{1, 2, 3, 4, 5})
			}
		})
	}
}
