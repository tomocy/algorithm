package main

func main() {}

func solve(vs []int, target int) bool {
	min, max := 0, len(vs)
	for min <= max {
		mid := (min + max) / 2
		v := vs[mid]
		if v == target {
			return true
		}
		if target < v {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return false
}
