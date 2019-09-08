package main

import "sort"

func main() {}

func simpleSearch(items []int, target int) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}

func binarySearch(items []int, target int) bool {
	sort.Ints(items)
	begin, end := 0, len(items)-1
	for begin <= end {
		mid := (begin + end) / 2
		guessed := items[mid]
		if guessed == target {
			return true
		}
		if target < guessed {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return false
}
