package main

func main() {}

func selectionSort(items []int) []int {
	sorteds := make([]int, len(items))
	copy(sorteds, items)
	for i := range sorteds {
		min, next := i, i+1
		for j := range sorteds[next:] {
			guessed := next + j
			if sorteds[guessed] < sorteds[min] {
				min = guessed
			}
		}

		temp := sorteds[i]
		sorteds[i], sorteds[min] = sorteds[min], temp
	}

	return sorteds
}
