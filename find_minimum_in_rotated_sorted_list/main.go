package main

func main() {}

func simple(vs []int) int {
	var lastV int
	for i, v := range vs {
		if i != 0 && v < lastV {
			return v
		}

		lastV = v
	}

	return vs[0]
}

func binarySearch(vs []int) int {
	begin, end := 0, len(vs)-1
	if vs[begin] < vs[end] {
		return vs[begin]
	}

	for begin <= end {
		mid := (begin + end) / 2
		left, right := mid-1, mid+1

		if left < begin {
			min := vs[mid]
			if right <= end && vs[right] < min {
				min = vs[right]
			}
			return min
		}
		if end < right {
			min := vs[mid]
			if begin <= left && vs[left] < min {
				min = vs[left]
			}
			return min
		}

		if vs[mid] < vs[left] {
			return vs[mid]
		}
		if vs[right] < vs[mid] {
			return vs[right]
		}

		if vs[begin] <= vs[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
