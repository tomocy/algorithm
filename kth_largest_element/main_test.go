package main

import "testing"

var solutions = map[string]func([]int, int) int{
	"with quick sort": withQuickSort,
}

func TestSolve(t *testing.T) {
	type input struct {
		ns []int
		k  int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{input{[]int{4, 5, 8, 2}, 3}, 4},
		{input{[]int{4, 5, 8, 2, 3}, 3}, 4},
		{input{[]int{4, 5, 8, 2, 3, 5}, 3}, 5},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.ns, test.input.k)
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
				s([]int{4, 5, 8, 2}, 3)
			}
		})
	}
}
