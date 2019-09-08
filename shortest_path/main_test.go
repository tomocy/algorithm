package main

import "testing"

var solutions = map[string]func(map[string][]string, string, string) int{}

func TestSolve(t *testing.T) {
	type input struct {
		graph         map[string][]string
		start, finish string
	}
	tests := []struct {
		input    input
		expected int
	}{
		{input{map[string][]string{
			"a": []string{"b", "c", "d"}, "b": []string{"f"},
			"c": []string{"e", "f"}, "d": []string{"g", "h"},
			"e": []string{}, "f": []string{}, "g": []string{}, "h": []string{}, "i": []string{},
		}, "a", "g"}, 2},
		{input{map[string][]string{
			"a": []string{"b", "c", "d"}, "b": []string{"f"},
			"c": []string{"e", "f"}, "d": []string{"g", "h"},
			"e": []string{}, "f": []string{}, "g": []string{}, "h": []string{}, "i": []string{},
		}, "a", "i"}, -1},
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
				s(map[string][]string{
					"a": []string{"b", "c", "d"}, "b": []string{"f"},
					"c": []string{"e", "f"}, "d": []string{"g", "h"},
					"e": []string{}, "f": []string{}, "g": []string{}, "h": []string{}, "i": []string{},
				}, "a", "g")
			}
		})
	}
}
