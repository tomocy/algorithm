package main

import (
	"sort"
	"testing"
)

var solutions = map[string]func(map[string][]string, []string) []string{
	"greedy": greedy,
}

func TestSolve(t *testing.T) {
	type input struct {
		covers  map[string][]string
		neededs []string
	}
	tests := []struct {
		input     input
		expecteds []string
	}{
		{input{map[string][]string{
			"a": []string{"1", "2", "3"}, "b": []string{"4", "1", "5"}, "c": []string{"6", "2", "7"},
			"d": []string{"2", "3"}, "e": []string{"7", "8"},
		}, []string{"1", "2", "3", "4", "5", "6", "7", "8"}}, []string{"a", "b", "c", "e"}},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actuals := s(test.input.covers, test.input.neededs)
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
				s(map[string][]string{
					"a": []string{"1", "2", "3"}, "b": []string{"4", "1", "5"}, "c": []string{"6", "2", "7"},
					"d": []string{"2", "3"}, "e": []string{"7", "8"},
				}, []string{"1", "2", "3", "4", "5", "6", "7", "8"})
			}
		})
	}
}
