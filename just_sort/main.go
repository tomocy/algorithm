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

func quickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	pivot, rests := items[0], items[1:]
	less, greater := make([]int, 0, len(rests)), make([]int, 0, len(rests))
	for _, x := range rests {
		if x <= pivot {
			less = append(less, x)
		} else {
			greater = append(greater, x)
		}
	}
	less, greater = append(quickSort(less), pivot), quickSort(greater)

	return append(less, greater...)
}
