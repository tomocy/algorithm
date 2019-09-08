package main

import (
	"sort"
	"testing"
)

var solutions = map[string]func(map[string]*trip, float64) []string{}

func TestSolve(t *testing.T) {
	type input struct {
		trips map[string]*trip
		limit float64
	}
	tests := []struct {
		input     input
		expecteds []string
	}{
		{input{map[string]*trip{
			"a": {"a", 7, 0.5}, "b": {"b", 6, 0.5}, "c": {"c", 9, 1}, "d": {"d", 9, 2}, "e": {"e", 8, 0.5},
		}, 2}, []string{"a", "c", "e"}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actuals := s(test.input.trips, test.input.limit)
				if len(actuals) != len(test.expecteds) {
					t.Fatalf("unexpected length: got %d, expect %d\n", len(actuals), len(test.expecteds))
				}
				sort.Strings(actuals)
				for i, expected := range test.expecteds {
					if actuals[i] != expected {
						t.Errorf("got %s, expect %s\n", actuals[i], expected)
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
				s(map[string]*trip{
					"a": {"a", 7, 0.5}, "b": {"b", 6, 0.5}, "c": {"c", 9, 1}, "d": {"d", 9, 2}, "e": {"e", 8, 0.5},
				}, 2)
			}
		})
	}
}
