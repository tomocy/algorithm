package main

func main() {}

func recursive(n int) int {
	if n <= 1 {
		return 1
	}

	return n * recursive(n-1)
}
