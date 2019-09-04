package main

import (
	"fmt"
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func([]int) []int{
	"quick": quick,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected []int
	}{
		{[]int{3, 4, 2, 1, 6, 8}, []int{1, 2, 3, 4, 6, 8}},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input)
				if len(actual) != len(test.expected) {
					t.Fatalf("unexpected length: got %d, expected %d\n", len(actual), len(test.expected))
				}
				for i, expected := range test.expected {
					if actual[i] != expected {
						algorithm.Reportln(t, fmt.Sprintf("value at %d", i), actual[i], expected)
					}
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{3, 4, 2, 1, 6, 8})
			}
		})
	}
}
