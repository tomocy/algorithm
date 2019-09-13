package main

func main() {}

func recursion(n, k int) int {
	if n == 1 {
		return 0
	}

	l := k
	if k%2 != 0 {
		l++
	}
	original := recursion(n-1, l/2)

	if k%2 != 0 {
		return original
	}
	return (original + 1) % 2
}
