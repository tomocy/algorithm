package main

func main() {}

func recursive(root *node) int {
	if root == nil {
		return 0
	}

	switch {
	case root.left != nil && root.right != nil:
		min := 1 + recursive(root.left)
		if guessed := 1 + recursive(root.right); guessed < min {
			min = guessed
		}
		return min
	case root.left == nil && root.right != nil:
		return 1 + recursive(root.right)
	case root.left != nil && root.right == nil:
		return 1 + recursive(root.left)
	default:
		return 1
	}
}

type node struct {
	val         int
	left, right *node
}
