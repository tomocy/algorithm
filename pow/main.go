package main

func main() {}

func recursion(x float64, n int) float64 {
	if n < 0 {
		return 1 / recursion(x, -n)
	}

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	r := recursion(x, n/2)
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}
}
