package main

func main() {}

func dijkstra(graph map[string]map[string]int, start, finish string) int {
	ns := parseGraph(graph)
	costs := make(costs)
	startN := ns[start]
	if startN.left != nil {
		costs[startN.left.node.val] = startN.left.cost
	}
	if startN.right != nil {
		costs[startN.right.node.val] = startN.right.cost
	}
	for n, cost := costs.min(ns); n != nil; n, cost = costs.min(ns) {
		if n.left != nil {
			costs[n.left.node.val] = costs.less(cost+n.left.cost, n.left.node.val)
		}
		if n.right != nil {
			costs[n.right.node.val] = costs.less(cost+n.right.cost, n.right.node.val)
		}

		ns[n.val].processed = true
	}

	return costs[finish]
}

func parseGraph(graph map[string]map[string]int) map[string]*node {
	ns := make(map[string]*node)
	for parent, children := range graph {
		if ns[parent] == nil {
			ns[parent] = &node{
				val: parent,
			}
		}
		var i int
		for child, cost := range children {
			if ns[child] == nil {
				ns[child] = &node{
					val: child,
				}
			}
			e := &edge{
				cost: cost, node: ns[child],
			}
			if i == 0 {
				ns[parent].left = e
			} else {
				ns[parent].right = e
			}
			i++
		}
	}

	return ns
}

type edge struct {
	cost int
	node *node
}

type node struct {
	val         string
	left, right *edge
	processed   bool
}

type costs map[string]int

func (cs costs) min(ns map[string]*node) (*node, int) {
	var minCost int
	var minNode *node
	for name, cost := range cs {
		n := ns[name]
		if n.processed {
			continue
		}
		if minNode == nil {
			minCost = cost
			minNode = n
		} else if cost < minCost {
			minCost = cost
			minNode = n
		}
	}

	return minNode, minCost
}

func (cs costs) less(a int, target string) int {
	less := a
	if stored, ok := cs[target]; ok && stored < less {
		less = stored
	}

	return less
}
