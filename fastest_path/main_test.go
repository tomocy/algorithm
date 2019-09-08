package main

import "testing"

var solutions = map[string]func(map[string]map[string]int, string, string) int{
	"dyikstra": dyikstra,
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
		{input{map[string]map[string]int{
			"a": {"b": 6, "c": 2},
			"b": {"d": 1},
			"c": {"b": 3, "d": 5},
		}, "a", "d"}, 6},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.graph, test.input.start, test.input.finish)
				if actual != test.expected {
					t.Errorf("got %d, expect %d\n", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s(map[string]map[string]int{
					"start": {"a": 6, "b": 2},
					"a":     {"finish": 1},
					"b":     {"a": 3, "finish": 5},
				}, "start", "finish")
			}
		})
	}
}
