package main

func main() {}

func twoPointers(head *node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow, fast := head, head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}

		slow, fast = slow.next, fast.next.next
	}

	return true
}

type node struct {
	val  int
	next *node
}
