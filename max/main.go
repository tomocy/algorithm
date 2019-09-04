package main

func main() {}

func selection(vs []int) int {
	var max int
	for i, v := range vs {
		if i == 0 {
			max = v
			continue
		}

		if max < v {
			max = v
		}
	}

	return max
}
