package main

func main() {}

func solve(vs []int) []int {
	if len(vs) < 2 {
		return vs
	}

	pivot := vs[0]
	var less, greater []int
	for _, v := range vs[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less, right := append(solve(less), pivot), solve(greater)

	return append(less, right...)
}
