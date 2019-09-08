package main

import (
	"sort"
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func(map[string][]string, []string) []string{
	"greedy": greedy,
}

func TestSolve(t *testing.T) {
	type input struct {
		nodes  map[string][]string
		needed []string
	}
	tests := []struct {
		input    input
		expected []string
	}{
		{
			input{
				map[string][]string{
					"k1": []string{"id", "nv", "ut"},
					"k2": []string{"wa", "id", "mt"},
					"k3": []string{"or", "nv", "ca"},
					"k4": []string{"nv", "ut"},
					"k5": []string{"ca", "az"},
				},
				[]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"},
			},
			[]string{"k1", "k2", "k3", "k5"},
		},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.nodes, test.input.needed)
				if len(actual) != len(test.expected) {
					t.Fatalf("unexpected length: got %d, expected %d\n", len(actual), len(test.expected))
				}
				sort.Strings(actual)
				for i, expected := range test.expected {
					if actual[i] != expected {
						algorithm.Reportln(t, "best cover", actual[i], expected)
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
				s(map[string][]string{
					"k1": []string{"id", "nv", "ut"},
					"k2": []string{"wa", "id", "mt"},
					"k3": []string{"or", "nv", "ca"},
					"k4": []string{"nv", "ut"},
					"k5": []string{"ca", "az"},
				}, []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
			}
		})
	}
}
