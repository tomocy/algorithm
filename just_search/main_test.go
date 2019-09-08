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
		{input{[]int{3, 4, 1, 6, 7, 9, 0}, 7}, true},
		{input{[]int{3, 4, 1, 6, 7, 9, 0}, -1}, false},
		{input{[]int{}, 0}, false},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input.items, test.input.target)
				if actual != test.expected {
					t.Errorf("got %t, expect %t\n", actual, test.expected)
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		for i := 0; i < b.N; i++ {
			b.Run(name, func(b *testing.B) {
				s([]int{3, 4, 1, 6, 7, 9, 0}, 7)
			})
		}
	}
}
