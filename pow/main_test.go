package main

import "testing"

var solutions = map[string]func(float64, int) float64{}

func TestSolve(t *testing.T) {
	type input struct {
		x float64
		n int
	}
	tests := []struct {
		input    input
		expected float64
	}{
		{input{2, 10}, 1024},
		{input{2.1, 3}, 9.261},
		{input{2, -2}, 0.25},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.x, test.input.n)
				if actual != test.expected {
					t.Errorf("got %F, expect %F", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(2, 10)
			}
		})
	}
}
