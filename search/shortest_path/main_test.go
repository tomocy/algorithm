package main

import (
	"testing"

	"github.com/tomocy/algorithm"
)

var solutions = map[string]func(map[string][]string, string) int{
	"breadth first search": breadthFirstSearch,
}

func TestSolve(t *testing.T) {
	type input struct {
		graph map[string][]string
		start string
	}
	tests := []struct {
		input    input
		expected int
	}{
		{
			input{
				map[string][]string{
					"you":    []string{"alice", "bob", "claire"},
					"bob":    []string{"anuj", "peggy"},
					"alice":  []string{"peggy"},
					"claire": []string{"thom", "jonny"},
					"anuj":   []string{}, "peggy": []string{}, "thom": []string{}, "jonny": []string{},
				},
				"you",
			},
			2,
		},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.graph, test.input.start)
				if actual != test.expected {
					algorithm.Reportln(t, "shortest path", actual, test.expected)
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
					"you":    []string{"alice", "bob", "claire"},
					"bob":    []string{"anuj", "peggy"},
					"alice":  []string{"peggy"},
					"claire": []string{"thom", "jonny"},
					"anuj":   []string{}, "peggy": []string{}, "thom": []string{}, "jonny": []string{},
				}, "you")
			}
		})
	}
}
