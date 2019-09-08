package main

func main() {}

func simple(items []int) int {
	var max int
	for i, item := range items {
		if i == 0 || max < item {
			max = item
		}
	}

	return max
}

func recursive(items []int) int {
	switch len(items) {
	case 0:
		return 0
	case 1:
		return items[0]
	default:
		max := items[0]
		guessed := recursive(items[1:])
		if max < guessed {
			max = guessed
		}

		return max
	}
}
