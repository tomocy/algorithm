package main

import (
	"log"
	"sort"
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func(map[string]item, float64) []string{
	"dynamic": dynamic,
}

func TestSolve(t *testing.T) {
	type input struct {
		items map[string]item
		limit float64
	}
	tests := []struct {
		input    input
		expected []string
	}{
		{
			input{
				map[string]item{
					"stereo": {3000, 4}, "laptop": {2000, 3}, "guiter": {1500, 1},
				},
				4,
			},
			[]string{"guiter", "laptop"},
		},
	}

	for _, test := range tests {
		for name, s := range solutions {
			t.Run(name, func(t *testing.T) {
				actual := s(test.input.items, test.input.limit)
				if len(actual) != len(test.expected) {
					t.Fatalf("unexpected length: got %d, expect %d\n", len(actual), len(test.expected))
				}
				sort.Strings(actual)
				log.Println(actual)
				for i, expected := range test.expected {
					if actual[i] != expected {
						algorithm.Reportln(t, "knapsack", actual[i], expected)
					}
				}
			})
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			s(map[string]item{
				"stereo": {3000, 4}, "laptop": {2000, 3}, "guiter": {1500, 1},
			}, 4)
		})
	}
}
