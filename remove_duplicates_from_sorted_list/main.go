package main

func main() {}

func simple(head *node) *node {
	current := head
	for current != nil && current.next != nil {
		if current.val == current.next.val {
			current.next = current.next.next
			continue
		}

		current = current.next
	}

	return head
}

type node struct {
	val  int
	next *node
}
