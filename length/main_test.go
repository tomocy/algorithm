package main

import "testing"

var solutions = map[string]func([]interface{}) int{
	"builtin": builtin,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		inputs   []interface{}
		expected int
	}{
		{[]interface{}{1, "1", true}, 3},
		{[]interface{}{}, 0},
	}

	for name, s := range solutions {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				actual := s(test.inputs)
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
				s([]interface{}{1, "1", true})
			}
		})
	}
}
