package main

import (
	"testing"
)

var solutions = map[string]func([]int) int{
	"simple": simple, "solve2": solve2,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 15},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
				if actual != test.expected {
					t.Errorf("got %d, expect %d\n", actual, test.expected)
				}
			}
		})
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
