package main

import "sort"

func main() {}

func greedy(ns map[string][]string, needed []string) []string {
	var bests []string
	ks := nodes(ns).sortKeys()
	for needed := newSet(needed); 0 < len(needed); {
		var best string
		var covered set
		for _, name := range ks {
			n := ns[name]
			copied := make(set, len(needed))
			copy(copied, needed)
			copied.intersect(newSet(n))
			if len(covered) < len(copied) {
				best = name
				covered = copied
			}
		}

		bests = append(bests, best)
		needed.diff(covered)
	}

	return bests
}

type nodes map[string][]string

func (ns nodes) sortKeys() []string {
	ks := make([]string, len(ns))
	var i int
	for k := range ns {
		ks[i] = k
		i++
	}
	sort.Strings(ks)

	return ks
}

func newSet(xs []string) set {
	var set set
	set.add(xs...)
	return set
}

type set []string

func (s *set) diff(other set) {
	var diffed set
	for _, item := range *s {
		if !other.isIn(item) {
			diffed.add(item)
		}
	}

	*s = diffed
}

func (s *set) intersect(other set) {
	var intersected set
	for _, item := range *s {
		if other.isIn(item) {
			intersected.add(item)
		}
	}

	*s = intersected
}

func (s *set) add(xs ...string) {
	for _, x := range xs {
		if !(*s).isIn(x) {
			*s = append(*s, x)
		}
	}
}

func (s set) isIn(target string) bool {
	for _, item := range s {
		if item == target {
			return true
		}
	}

	return false
}
