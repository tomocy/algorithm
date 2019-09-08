package main

func main() {}

func builtin(items []interface{}) int {
	return len(items)
}

func recursive(items []interface{}) (l int) {
	defer func() {
		if err := recover(); err != nil {
			l = 0
		}
	}()

	return 1 + recursive(items[1:])
}
