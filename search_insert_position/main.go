package main

func main() {}

func binarySearch(vs []int, x int) int {
	begin, end := 0, len(vs)-1
	for begin <= end {
		mid := (begin + end) / 2
		if vs[mid] == x {
			return mid
		}
		if vs[mid] < x {
			begin = mid + 1
		} else {
			end = mid - 1
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
