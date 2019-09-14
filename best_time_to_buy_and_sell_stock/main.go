package main

func main() {}

func dynamic(ps []int) int {
	grid := make(grid)
	var lastP, lastQ int
	for i, p := range ps {
		lastQ = 0
		for j, q := range ps {
			max := grid.cell(p, lastQ)
			if guessed := grid.cell(lastP, q); max < guessed {
				max = guessed
			}
			if i <= j {
				if guessed := cell(q - p); max < guessed {
					max = guessed
				}
			}

			grid.setCell(p, q, max)
			lastQ = q
		}
		lastP = p
	}

	return int(grid.cell(lastP, lastQ))
}

type grid map[int]map[int]cell

func (g grid) cell(i, j int) cell {
	if g[i] == nil {
		return 0
	}

	return g[i][j]
}

func (g *grid) setCell(i, j int, c cell) {
	if (*g)[i] == nil {
		(*g)[i] = make(map[int]cell)
	}

	(*g)[i][j] = c
}

type cell int

func maxFromMin(ps []int) int {
	return -1
}
