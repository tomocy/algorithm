package main

import "testing"

var solutions = map[string]func([]string, string) string{
	"dynamic": dynamic,
}

func TestSolve(t *testing.T) {
	type input struct {
		vocabs []string
		s      string
	}
	tests := []struct {
		input    input
		expected string
	}{
		{input{[]string{"vista", "fish"}, "hish"}, "fish"},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.input.vocabs, test.input.s)
				if actual != test.expected {
					t.Errorf("got %s, expect %s\n", actual, test.expected)
				}
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for name, s := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s([]string{"vista", "fish"}, "hish")
			}
		})
	}
}
