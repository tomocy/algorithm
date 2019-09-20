package main

func main() {}

func recursive(root *node) int {
	if root == nil {
		return 0
	}

	max := 1 + recursive(root.left)
	if guessed := 1 + recursive(root.right); max < guessed {
		max = guessed
	}

	return max
}

type node struct {
	val         int
	left, right *node
}
