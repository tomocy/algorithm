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
	return -1
}
