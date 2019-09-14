package main

func main() {}

func dynamic(ps []int) int {
	graph := make(graph)
	var lastP, lastQ int
	for i, p := range ps {
		lastQ = 0
		for j, q := range ps {
			max := graph.cell(p, lastQ)
			if guessed := graph.cell(lastP, q); max < guessed {
				max = guessed
			}
			if i <= j {
				if guessed := cell(q - p); max < guessed {
					max = guessed
				}
			}

			graph.setCell(p, q, max)
			lastQ = q
		}

		lastP = p
	}

	return int(graph.cell(lastP, lastQ))
}

type graph map[int]map[int]cell

func (g graph) cell(i, j int) cell {
	if g[i] == nil {
		return 0
	}

	return g[i][j]
}

func (g *graph) setCell(i, j int, c cell) {
	if (*g)[i] == nil {
		(*g)[i] = make(map[int]cell)
	}

	(*g)[i][j] = c
}

type cell int
