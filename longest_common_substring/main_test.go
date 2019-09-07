package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func([]string, string) string{
	"dynamic": dynamic,
}

func TestSolve(t *testing.T) {
	type input struct {
		items []string
		name  string
	}
	tests := []struct {
		input    input
		expected string
	}{
		{
			input{
				[]string{"vista", "fish"},
				"hish",
			},
			"fish",
		},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input.items, test.input.name)
				if actual != test.expected {
					algorithm.Reportln(t, "longest common substring", actual, test.expected)
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			s([]string{"fish", "vista"}, "hish")
		})
	}
}
