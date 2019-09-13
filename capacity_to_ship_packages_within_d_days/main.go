package main

func main() {}

func binarySearch(ps []int, d int) int {
	max := func(vs []int) int {
		var max int
		for i, v := range vs {
			if i == 0 || max < v {
				max = v
			}
		}

		return max
	}
	sum := func(vs []int) int {
		var sum int
		for _, v := range vs {
			sum += v
		}

		return sum
	}

	begin, end := max(ps), sum(ps)
	for begin < end {
		cap := (begin + end) / 2
		var sum int
		neededs := 1
		for _, p := range ps {
			if cap < sum+p {
				neededs++
				sum = p
				continue
			}

			sum += p
		}

		if d < neededs {
			begin = cap + 1
		} else {
			end = cap
		}
	}

	return begin
}
