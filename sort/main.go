package main

func main() {}

func quick(vs []int) []int {
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
	less, greater = quick(less), quick(greater)

	return append(append(less, pivot), greater...)
}

func selection(vs []int) []int {
	for i := range vs {
		min := i
		next := i + 1
		for j := range vs[next:] {
			idx := next + j
			if vs[idx] < vs[min] {
				min = idx
			}
		}

		temp := vs[i]
		vs[i], vs[min] = vs[min], temp
	}

	return vs
}
