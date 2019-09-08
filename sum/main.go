package main

func main() {}

func simple(items []int) int {
	var sum int
	for _, item := range items {
		sum += item
	}

	return sum
}
