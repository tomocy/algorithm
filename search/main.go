package main

func main() {}

func simpleSearch(items []int, target int) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}
