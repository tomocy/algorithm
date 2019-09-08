package main

import (
	"sort"
	"testing"
)

var solutions = map[string]func(map[string]*item, float64) []string{}

func TestSolve(t *testing.T) {
	type input struct {
		items map[string]*item
		limit float64
	}
	tests := []struct {
		input     input
		expecteds []string
	}{
		{input{map[string]*item{
			"a": {"a", 3000, 4}, "b": {"b", 2000, 3}, "c": {"c", 1500, 1},
		}, 4}, []string{"b", "c"}},
		{input{map[string]*item{
			"a": {"a", 10, 3}, "b": {"c", 3, 1}, "c": {"c", 9, 2}, "d": {"d", 5, 2}, "e": {"e", 6, 1},
		}, 6}, []string{"a", "c", "e"}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actuals := s(test.input.items, test.input.limit)
				if len(actuals) != len(test.expecteds) {
					t.Fatalf("unexpected langth: got %d, expect %d\n", len(actuals), len(test.expecteds))
				}
				sort.Strings(actuals)
				for i, expected := range test.expecteds {
					if actuals[i] != expected {
						t.Errorf("got %s, expect %s", actuals[i], expected)
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
				s(map[string]*item{
					"a": {"a", 3000, 4}, "b": {"b", 2000, 3}, "c": {"c", 1500, 1},
				}, 4)
			}
		})
	}
}
