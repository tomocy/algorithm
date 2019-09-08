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

func quickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	pivot, rests := items[0], items[1:]
	less, greater := make([]int, 0, len(items)), make([]int, 0, len(rests))
	for _, item := range rests {
		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}
	less, greater = append(quickSort(less), pivot), quickSort(greater)

	return append(less, greater...)
}
