package main

func main() {}

func simple(vs []int) int {
	if len(vs) < 1 {
		return 0
	}
	lastV := vs[0]
	for _, v := range vs[1:] {
		if v < lastV {
			return v
		}
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
		if left < 0 {
			min := vs[mid]
			if right < len(vs) && vs[right] < min {
				min = vs[right]
			}
			return min
		}
		if len(vs) <= right {
			min := vs[mid]
			if 0 <= left && vs[left] < min {
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

		if vs[0] < vs[mid] {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
