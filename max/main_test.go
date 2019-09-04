package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func([]int) int{
	"selection": selection,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{4, 5, 2, 3, 5, 6, 8}, 8},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input)
				if actual != test.expected {
					algorithm.Reportln(t, "max", actual, test.expected)
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{4, 5, 2, 3, 5, 6, 8})
			}
		})
	}
}
