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
