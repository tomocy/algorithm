package main

func main() {}

func solve(vs []int, target int) bool {
	for _, v := range vs {
		if v == target {
			return true
		}
	}

	return false
}
