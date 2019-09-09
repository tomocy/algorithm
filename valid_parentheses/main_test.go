package main

import "testing"

var solutions = map[string]func(string) bool{}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true}, {"()[]{}", true}, {"([)]", false}, {"{[]}", true},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
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
				s("()")
			}
		})
	}
}
