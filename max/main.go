package main

func main() {}

func simple(vs []int) int {
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

func recursive(vs []int) int {
	switch len(vs) {
	case 0:
		return 0
	case 1:
		return vs[0]
	default:
		max := vs[0]
		alt := recursive(vs[1:])
		if max < alt {
			max = alt
		}

		return max
	}
}
