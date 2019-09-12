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
	return -1
}
