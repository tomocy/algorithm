package main

func main() {}

func solve(vs []int) int {
	return len(vs)
}

func solve2(vs []int) (l int) {
	defer func() {
		if err := recover(); err != nil {
			l = 0
		}
	}()

	return 1 + solve(vs[1:])
}
