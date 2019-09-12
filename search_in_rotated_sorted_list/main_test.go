package main

import "testing"

var solutions = map[string]func([]int, int) int{
	"simple": simple,
}

func TestSolve(t *testing.T) {
	type input struct {
		vs []int
		x  int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{input{[]int{4}, 4}, 0},
		{input{[]int{4, 5}, 3}, -1},
		{input{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{input{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.vs, test.input.x)
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
				s([]int{4, 5, 6, 7, 0, 1, 2}, 0)
			}
		})
	}
}
