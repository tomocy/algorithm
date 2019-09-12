package main

func main() {}

func binarySearch(vs []int, x int) int {
	begin, end := 0, len(vs)-1
	if vs[end] < x {
		return end + 1
	}

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

	return begin
}
