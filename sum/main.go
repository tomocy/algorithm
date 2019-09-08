package main

func main() {}

func simple(vs []int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}

	return sum
}

func recursive(vs []int) int {
	if len(vs) < 1 {
		return 0
	}

	return vs[0] + recursive(vs[1:])
}
