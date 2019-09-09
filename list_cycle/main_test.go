package main

import "testing"

var solutions = map[string]func(head *node) bool{}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    func() *node
		expected bool
	}{
		{func() *node {
			a, b, c, d := &node{val: 3}, &node{val: 2}, &node{val: 0}, &node{val: -4}
			a.next, b.next, c.next, d.next = b, c, d, b
			return a
		}, true},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input())
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
				s(func() *node {
					a, b, c, d := &node{val: 3}, &node{val: 2}, &node{val: 0}, &node{val: -4}
					a.next, b.next, c.next, d.next = b, c, d, b
					return a
				}())
			}
		})
	}
}
