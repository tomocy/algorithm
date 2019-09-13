package main

import "testing"

var solutions = map[string]func([]int, int) int{}

func TestSolve(t *testing.T) {
	type input struct {
		ps []int
		d  int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{input{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, 15},
		{input{[]int{3, 2, 2, 4, 1, 4}, 3}, 6},
		{input{[]int{1, 2, 3, 1, 1}, 4}, 3},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.ws, test.input.d)
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
				s([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
			}
		})
	}
}
