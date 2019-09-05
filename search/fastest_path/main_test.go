package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func(map[string]map[string]int, string, string) int{
	"dijkstra": dijkstra,
}

func TestSolve(t *testing.T) {
	type input struct {
		graph         map[string]map[string]int
		start, finish string
	}
	tests := []struct {
		input    input
		expected int
	}{
		{
			input{
				map[string]map[string]int{
					"start": {"a": 6, "b": 2},
					"a":     {"finish": 1},
					"b":     {"a": 3, "finish": 5},
				},
				"start", "finish",
			},
			6,
		},
		{
			input{
				map[string]map[string]int{
					"start": {"a": 5, "b": 2},
					"a":     {"c": 4, "d": 2},
					"b":     {"a": 8, "d": 7},
					"c":     {"finish": 3, "d": 6},
					"d":     {"finish": 1},
				},
				"start", "finish",
			},
			8,
		},
		{
			input{
				map[string]map[string]int{
					"start": {"a": 10},
					"a":     {"b": 20},
					"b":     {"c": 1, "finish": 30},
					"c":     {"a": 1},
				},
				"start", "finish",
			},
			60,
		},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input.graph, test.input.start, test.input.finish)
				if actual != test.expected {
					algorithm.Reportln(t, "fastest path", actual, test.expected)
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			s(map[string]map[string]int{
				"start": {"a": 6, "b": 2},
				"a":     {"finish": 1},
				"b":     {"a": 3, "finish": 5},
			}, "start", "finish")
		})
	}
}
