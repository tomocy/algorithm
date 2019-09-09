package main

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

type node struct {
	val  int
	next *node
}
