package main

func main() {}

func selectionSort(items []int) []int {
	sorteds := make([]int, len(items))
	copy(sorteds, items)
	for i := range sorteds {
		min := i
		next := i + 1
		for j := range sorteds[next:] {
			x := next + j
			if sorteds[x] < sorteds[min] {
				min = x
			}
		}

		temp := sorteds[i]
		sorteds[i], sorteds[min] = sorteds[min], temp
	}

	return sorteds
}
