package main

import "testing"

var solutions = map[string]func(*node) int{}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    *node
		expected int
	}{
		{
			input:    &node{val: 3, left: &node{val: 9}, right: &node{val: 20, left: &node{val: 15}, right: &node{val: 7}}},
			expected: 3,
		},
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
				s(&node{val: 3, left: &node{val: 9}, right: &node{val: 20, left: &node{val: 15}, right: &node{val: 7}}})
			}
		})
	}
}
