package main

func main() {}

func solve(vs []int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}

	return sum
}
