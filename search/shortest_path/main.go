package main

import (
	"strings"
)

func main() {}

func breadthFirstSearch(graph map[string][]string, start string) int {
	ns := parseGraph(graph)
	q := queue{start}
	ns[start].queued = true
	for v := q.dequeue(); v != ""; v = q.dequeue() {
		neighors := graph[v]
		for _, neighor := range neighors {
			n := ns[neighor]
			n.paths = ns[v].paths + 1
			if n.isTarget() {
				return n.paths
			}
			if n.queued {
				continue
			}

			q = append(q, neighor)
			ns[neighor].queued = true
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
	val    string
	paths  int
	queued bool
}

func (n *node) isTarget() bool {
	return strings.HasSuffix(n.val, "m")
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
