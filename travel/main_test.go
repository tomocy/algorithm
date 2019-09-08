package main

import (
	"sort"
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func(map[string]tour, float64) []string{
	"dynamic": dynamic,
}

func TestSolve(t *testing.T) {
	type input struct {
		items map[string]tour
		limit float64
	}
	tests := []struct {
		input    input
		expected []string
	}{
		{
			input{
				map[string]tour{
					"a": {7, 0.5}, "b": {6, 0.5}, "c": {9, 1}, "d": {9, 2}, "e": {8, 0.5},
				},
				2,
			},
			[]string{"a", "c", "e"},
		},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.items, test.input.limit)
				if len(actual) != len(test.expected) {
					t.Fatalf("unexpected length: got %d, expect %d\n", len(actual), len(test.expected))
				}
				sort.Strings(actual)
				for i, expected := range test.expected {
					if actual[i] != expected {
						algorithm.Reportln(t, "travel", actual[i], expected)
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
				s(map[string]tour{
					"a": {7, 0.5}, "b": {6, 0.5}, "c": {9, 1}, "d": {9, 2}, "e": {8, 0.5},
				}, 2)
			}
		})
	}
}
