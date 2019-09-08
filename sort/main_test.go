package main

import "testing"

var solutions = map[string]func([]int) []int{}

func TestSolve(t *testing.T) {
	tests := []struct {
		inputs, expecteds []int
	}{
		{[]int{2, 7, 2, 3, 0, 9, -1, 0}, []int{-1, 0, 0, 2, 2, 3, 7, 9}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{}, []int{}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actuals := s(test.inputs)
				if len(actuals) != len(test.expecteds) {
					t.Fatalf("unexpected length: got %d, expect %d\n", len(actuals), len(test.expecteds))
				}
				for i, expected := range test.expecteds {
					if actuals[i] != expected {
						t.Errorf("got %d, expect %d\n", actuals[i], expected)
					}
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]int{2, 7, 2, 3, 0, 9, -1, 0})
			}
		})
	}
}
