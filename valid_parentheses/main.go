package main

import (
	"strings"
)

func main() {}

var parens = map[rune]rune{
	'(': ')', '{': '}', '[': ']',
}

func simple(s string) bool {
	if len(s) == 1 {
		return false
	}

	var stack stack
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.push(r)
		case ')', '}', ']':
			popped := stack.pop()
			if r != parens[popped] {
				return false
			}
		}
	}

	return len(stack) == 0
}

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() rune {
	if len(*s) < 1 {
		return 0
	}

	i := len(*s) - 1
	last := (*s)[i]
	*s = (*s)[:i]

	return last
}

func (s stack) String() string {
	var b strings.Builder
	b.WriteRune('[')
	for i, r := range s {
		if i != 0 {
			b.WriteRune(' ')
		}

		b.WriteString(string(r))
	}
	b.WriteRune(']')

	return b.String()
}
