package main

func main() {}

func solve(n int) int {
	if n <= 1 {
		return 1
	}

	return n * solve(n-1)
}
