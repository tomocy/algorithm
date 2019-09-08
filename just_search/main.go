package main

import "sort"

func main() {}

func binarySearch(items []int, target int) bool {
	sort.Ints(items)
	begin, end := 0, len(items)-1
	for begin <= end {
		mid := (begin + end) / 2
		x := items[mid]
		if x == target {
			return true
		}
		if target < x {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return false
}
