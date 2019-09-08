package main

func main() {}

func greedy(covers map[string][]string, neededs []string) []string {
	var bests []string
	neededSet := newSet(neededs)
	for 0 < len(neededSet) {
		var best string
		var bestSet set
		for cover, covereds := range covers {
			guesseds := neededSet.intersect(newSet(covereds))
			if len(bestSet) < len(guesseds) {
				best, bestSet = cover, guesseds
			}
		}

		bests, neededSet = append(bests, best), neededSet.diff(bestSet)
	}

	return bests
}

func newSet(ss []string) set {
	set := make(set, 0, len(ss))
	for _, s := range ss {
		if !set.isIn(s) {
			set = append(set, s)
		}
	}

	return set
}

type set []string

func (s set) intersect(other set) set {
	var intersected set
	for _, e := range s {
		if other.isIn(e) {
			intersected = append(intersected, e)
		}
	}

	return intersected
}

func (s set) diff(other set) set {
	var diffed set
	for _, e := range s {
		if !other.isIn(e) {
			diffed = append(diffed, e)
		}
	}

	return diffed
}

func (s set) isIn(target string) bool {
	for _, e := range s {
		if e == target {
			return true
		}
	}

	return false
}
