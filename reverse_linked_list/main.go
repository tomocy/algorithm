package main

import (
	"fmt"
)

func main() {}

func simple(head *node) *node {
	var prev *node
	current := head
	for current != nil {
		temp := current.next
		current.next = prev
		prev, current = current, temp
	}

	return prev
}

func withStack(head *node) *node {
	var stack stack
	for n := head; n != nil; n = n.next {
		stack.push(n.val)
	}
	if len(stack) < 1 {
		return nil
	}

	popped, _ := stack.pop()
	reversed := &node{
		val: popped,
	}
	current := reversed
	for v, ok := stack.pop(); ok; v, ok = stack.pop() {
		current.next = &node{
			val: v,
		}
		current = current.next
	}

	return reversed
}

type stack []int

func (s *stack) push(a int) {
	*s = append(*s, a)
}

func (s *stack) pop() (int, bool) {
	if len(*s) < 1 {
		return 0, false
	}

	i := len(*s) - 1
	last := (*s)[i]
	*s = (*s)[:i]

	return last, true
}

type node struct {
	val  int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d->%s", n.val, n.next)
}
