package main

import "testing"

var solutions = map[string]func([]int) []int{
	"selection sort": selectionSort, "quick sort": quickSort,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		inputs, expecteds []int
	}{
		{[]int{3, 4, 2, 4, 6, 0, 1, -3, -8}, []int{-8, -3, 0, 1, 2, 3, 4, 4, 6}},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actuals := s(test.inputs)
				if len(actuals) != len(test.expecteds) {
					t.Fatalf("unexpected length: got %d, expected %d\n", len(actuals), len(test.expecteds))
				}
				for i, expected := range test.expecteds {
					if actuals[i] != expected {
						t.Errorf("got %d, expected %d\n", actuals[i], expected)
					}
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			s([]int{3, 4, 2, 4, 6, 0, 1, -3, -8})
		})
	}
}
