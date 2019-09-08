package main

import "testing"

var solutions = map[string]func([]int, int) bool{
	"simple search": simpleSearch, "binary search": binarySearch,
}

func TestSolve(t *testing.T) {
	type input struct {
		items  []int
		target int
	}
	tests := []struct {
		input    input
		expected bool
	}{
		{input{[]int{2, 7, 2, 3, 0, 9, -1, 0}, 7}, true},
		{input{[]int{2, 7, 2, 3, 0, 9, -1, 0}, -7}, false},
		{input{[]int{}, 0}, false},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.items, test.input.target)
				if actual != test.expected {
					t.Errorf("got %t, expect %t\n", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{2, 7, 2, 3, 0, 9, -1, 0}, 7)
			}
		})
	}
}
