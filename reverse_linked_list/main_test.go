package main

import (
	"fmt"
	"testing"
)

var solutions = map[string]func(head *node) *node{
	"simple": simple,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected *node
	}{
		{&node{1, &node{2, &node{3, &node{4, &node{5, nil}}}}}, &node{5, &node{4, &node{3, &node{2, &node{1, nil}}}}}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
				if err := assertNode(actual, test.expected); err != nil {
					t.Errorf("%s\n", err)
				}
			}
		})
	}
}

func assertNode(actual, expected *node) error {
	if actual == nil && expected == nil {
		return nil
	}
	if actual == nil || expected == nil {
		return fmt.Errorf("unexpected node: got %v, expect %v", actual, expected)
	}
	if actual.val != expected.val {
		return fmt.Errorf("unexpected val of node: got %d, expect %d", actual.val, expected.val)
	}

	return assertNode(actual.next, expected.next)
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(&node{1, &node{2, &node{3, &node{4, &node{5, nil}}}}})
			}
		})
	}
}
