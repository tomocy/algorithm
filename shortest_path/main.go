package main

func main() {}

func breadthFirstSearch(graph map[string][]string, start, finish string) int {
	ns := parseGraph(graph)
	q := queue{start}
	ns[start].queued = true
	for val := q.dequeue(); val != ""; val = q.dequeue() {
		neighors := graph[val]
		for _, neighor := range neighors {
			ns[neighor].distance = ns[val].distance + 1
			if neighor == finish {
				return ns[neighor].distance
			}
			if ns[neighor].queued {
				continue
			}

			q, ns[neighor].queued = append(q, neighor), true
		}
	}

	return -1
}

func parseGraph(graph map[string][]string) map[string]*node {
	ns := make(map[string]*node)
	for val := range graph {
		ns[val] = &node{
			val: val,
		}
	}

	return ns
}

type node struct {
	val      string
	queued   bool
	distance int
}

type queue []string

func (q *queue) dequeue() string {
	if len(*q) < 1 {
		return ""
	}

	first := (*q)[0]
	*q = (*q)[1:]

	return first
}
