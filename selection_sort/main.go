package main

func main() {}

func solve(vs []int) []int {
	sorted := make([]int, len(vs))
	copy(sorted, vs)
	for i := range sorted {
		min := i
		next := i + 1
		for j := range sorted[next:] {
			idx := next + j
			if sorted[idx] < sorted[min] {
				min = idx
			}
		}

		temp := sorted[i]
		sorted[i], sorted[min] = sorted[min], temp
	}

	return sorted
}
