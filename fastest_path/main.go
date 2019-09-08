package main

import (
	"fmt"
)

func main() {}

func dyikstra(graph map[string]map[string]int, start, finish string) int {
	ns, cs := parseGraph(graph), make(costs)
	startN := ns[start]
	if startN.left != nil {
		cs[startN.left.node.val] = startN.left.cost
	}
	if startN.right != nil {
		cs[startN.right.node.val] = startN.right.cost
	}
	for min, cost := cs.min(ns); min != nil; min, cost = cs.min(ns) {
		if min.left != nil {
			cs[min.left.node.val] = cs.less(cost+min.left.cost, min.left.node.val)
		}
		if min.right != nil {
			cs[min.right.node.val] = cs.less(cost+min.right.cost, min.right.node.val)
		}

		ns[min.val].processed = true
	}

	return cs[finish]
}

func parseGraph(graph map[string]map[string]int) map[string]*node {
	ns := make(map[string]*node)
	for val, neighors := range graph {
		if ns[val] == nil {
			ns[val] = &node{
				val: val,
			}
		}
		var i int
		for neighor, cost := range neighors {
			if ns[neighor] == nil {
				ns[neighor] = &node{
					val: neighor,
				}
			}
			e := &edge{
				cost: cost, node: ns[neighor],
			}
			if i == 0 {
				ns[val].left = e
			} else {
				ns[val].right = e
			}
			i++
		}
	}

	return ns
}

type costs map[string]int

func (cs costs) min(ns map[string]*node) (*node, int) {
	var minVal string
	var minCost int
	var i int
	for val, cost := range cs {
		if ns[val].processed {
			continue
		}
		if i == 0 || cost < minCost {
			minVal, minCost = val, cost
		}
		i++
	}

	return ns[minVal], minCost
}

func (cs costs) less(a int, b string) int {
	less := a
	if guessed, ok := cs[b]; ok && guessed < less {
		less = guessed
	}

	return less
}

type node struct {
	val         string
	left, right *edge
	processed   bool
}

func (n node) String() string {
	return fmt.Sprintf("%s(%s)(%s)", n.val, n.left, n.right)
}

type edge struct {
	cost int
	node *node
}

func (e edge) String() string {
	return fmt.Sprintf("- %d ->%s", e.cost, e.node)
}
