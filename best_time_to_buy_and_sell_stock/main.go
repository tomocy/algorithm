package main

func main() {}

func dynamic(ps []int) int {
	return -1
}

func maxFromMin(ps []int) int {
	var min, max int
	for i, p := range ps {
		if i == 0 || p < min {
			min = p
		} else if guessed := p - min; max < guessed {
			max = guessed
		}
	}

	return max
}
