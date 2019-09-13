package main

import "testing"

var solutions = map[string]func(int, int) int{}

func TestSolve(t *testing.T) {
	type input struct {
		n, k int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{input{1, 1}, 0},
		{input{2, 1}, 0},
		{input{2, 2}, 1},
		{input{4, 5}, 1},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.n, test.input.k)
				if actual != test.expected {
					t.Errorf("got %d, expect %d", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(1, 1)
			}
		})
	}
}
