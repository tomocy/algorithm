package main

import "testing"

var solutions = map[string]func([]int) int{
	"simple": simple, "recursive": recursive,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-1, 2}, 1},
		{[]int{-1}, -1},
		{[]int{}, 0},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.inputs)
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
				s([]int{1, 2, 3, 4, 5})
			}
		})
	}
}
