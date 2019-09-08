package main

func main() {}

func simple(items []int) int {
	var sum int
	for _, item := range items {
		sum += item
	}

	return sum
}

func recursive(items []int) int {
	if len(items) < 1 {
		return 0
	}

	return items[0] + recursive(items[1:])
}
