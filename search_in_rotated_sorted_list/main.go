package main

func main() {}

func simple(vs []int, x int) int {
	for i, v := range vs {
		if v == x {
			return i
		}
	}

	return -1
}

func binarySearch(vs []int, x int) int {
	begin, end := 0, len(vs)-1
	for begin <= end {
		mid := (begin + end) / 2
		if vs[mid] == x {
			return mid
		}

		if vs[begin] <= vs[mid] {
			if vs[begin] <= x && x < vs[mid] {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		} else {
			if vs[mid] < x && x <= vs[end] {
				begin = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
