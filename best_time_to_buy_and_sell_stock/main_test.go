package main

import "testing"

var solutions = map[string]func([]int) int{
	"dynamic": dynamic,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
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
				s([]int{7, 1, 5, 3, 6, 4})
			}
		})
	}
}
