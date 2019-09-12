package main

func main() {}

func simple(vs []int) int {
	if len(vs) < 1 {
		return 0
	}
	lastV := vs[0]
	for _, v := range vs[1:] {
		if v < lastV {
			return v
		}
	}

	return vs[0]
}
