package main

func main() {}

func simple(vs []int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}

	return sum
}

func solve2(vs []int) int {
	if len(vs) <= 0 {
		return 0
	}

	return vs[0] + solve2(vs[1:])
}
