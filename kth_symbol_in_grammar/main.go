package main

func main() {}

func recursion(n, k int) int {
	if n == 0 {
		return 0
	}

	l := k
	if k%2 != 0 {
		l++
	}

	x := recursion(n-1, l/2)
	if k%2 != 0 {
		return x
	}

	return (x + 1) % 2
}
