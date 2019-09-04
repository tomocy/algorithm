package main

import "math/rand"

func main() {}

func quick(vs []int) []int {
	if len(vs) < 2 {
		return vs
	}

	n := rand.Intn(len(vs))
	pivot := vs[n]
	var less, greater []int
	for _, v := range append(vs[:n], vs[n+1:]...) {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less, greater = append(quick(less), pivot), quick(greater)

	return append(less, greater...)
}
