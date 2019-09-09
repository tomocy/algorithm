package main

func main() {}

func withQuickSort(ns []int, k int) int {
	sorteds := quickSort(ns)
	return sorteds[len(ns)-k]
}

func quickSort(ns []int) []int {
	if len(ns) < 2 {
		return ns
	}

	pivot, rests := ns[0], ns[1:]
	less, greater := make([]int, 0, len(ns)), make([]int, 0, len(rests))
	for _, rest := range rests {
		if rest <= pivot {
			less = append(less, rest)
		} else {
			greater = append(greater, rest)
		}
	}
	less, greater = append(quickSort(less), pivot), quickSort(greater)

	return append(less, greater...)
}
