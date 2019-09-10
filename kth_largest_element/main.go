package main

func main() {}

func withQuickSort(ns []int, k int) int {
	sorteds := quickSort(ns)
	return sorteds[len(ns)-k]
}

func quickSort(ns []int) []int {
	if len(ns) < 2 {
		return ns
	}

	pivot, rests := ns[0], ns[1:]
	less, greater := make([]int, 0, len(ns)), make([]int, 0, len(rests))
	for _, rest := range rests {
		if rest <= pivot {
			less = append(less, rest)
		} else {
			greater = append(greater, rest)
		}
	}
	less, greater = append(quickSort(less), pivot), quickSort(greater)

	return append(less, greater...)
}

func withHeap(vs []int, k int) int {
	var heap heap
	for _, v := range vs {
		heap.push(v)
	}
	for i := 1; i < k; i++ {
		heap.pop()
	}

	return heap.max()
}

type heap []int

func (h *heap) push(v int) {
	*h = append(*h, v)
	pushedIdx := len(*h) - 1
	for 0 < pushedIdx {
		parentIdx := (pushedIdx - 1) / 2
		if v <= (*h)[parentIdx] {
			break
		}

		pushedIdx, (*h)[pushedIdx] = parentIdx, (*h)[parentIdx]
	}

	(*h)[pushedIdx] = v
}

func (h *heap) max() int {
	if len(*h) < 1 {
		return 0
	}

	return (*h)[0]
}

func (h *heap) pop() {
	if len(*h) < 1 {
		return
	}

	lastIdx := len(*h) - 1
	last := (*h)[lastIdx]
	*h = (*h)[:lastIdx]
	if len(*h) < 1 {
		return
	}

	var i int
	for leftIdx, rightIdx := i*2+1, i*2+2; leftIdx < len(*h); leftIdx, rightIdx = i*2+1, i*2+2 {
		maxIdx := leftIdx
		if rightIdx < len(*h) && (*h)[leftIdx] < (*h)[rightIdx] {
			maxIdx = rightIdx
		}
		if (*h)[maxIdx] <= last {
			break
		}

		i, (*h)[i] = maxIdx, (*h)[maxIdx]
	}

	(*h)[i] = last
}
