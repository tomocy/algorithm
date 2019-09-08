package main

func main() {}

func builtin(vs []int) int {
	return len(vs)
}

func recursive(vs []int) (l int) {
	defer func() {
		if err := recover(); err != nil {
			l = 0
		}
	}()

	return 1 + recursive(vs[1:])
}
