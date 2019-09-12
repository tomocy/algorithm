package main

func main() {}

func binarySearch(vs []int, x int) int {
	begin, end := 0, len(vs)-1
	for begin <= end {
		mid := (begin + end) / 2
		guessed := vs[mid]
		if guessed == x {
			return mid
		}
		if x < guessed {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	lastI := -1
	for i, v := range vs {
		if x < v {
			break
		}
		lastI = i
	}

	return lastI + 1
}
