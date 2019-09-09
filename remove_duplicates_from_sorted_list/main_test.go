package main

import (
	"fmt"
	"testing"
)

var solutions = map[string]func(*node) *node{
	"simple": simple,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input, expected *node
	}{
		{&node{val: 1, next: &node{val: 1, next: &node{val: 2}}}, &node{val: 1, next: &node{val: 2}}},
		{&node{val: 1, next: &node{val: 1, next: &node{val: 2, next: &node{val: 3, next: &node{val: 3}}}}}, &node{val: 1, next: &node{val: 2, next: &node{val: 3}}}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input)
				if err := assertNode(actual, test.expected); err != nil {
					t.Errorf("unexpected node: %s\n", err)
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
		return fmt.Errorf("unexpected val of node: got %v, expected %v", actual.val, expected.val)
	}

	return assertNode(actual.next, expected.next)
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(&node{val: 1, next: &node{val: 1, next: &node{val: 2}}})
			}
		})
	}
}
