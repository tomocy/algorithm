package main

func main() {}

func dynamic(vocabs []string, s string) string {
	var longest string
	var max int
	for _, vocab := range vocabs {
		guessed := countCommonSubstring(vocab, s)
		if max < guessed {
			longest, max = vocab, guessed
		}
	}

	return longest
}

func countCommonSubstring(a, b string) int {
	grid := newGrid(len(a), len(b))
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			continue
		}

		grid[i][i] = grid.cell(i-1, i-1) + 1
	}

	return grid.max()
}

func newGrid(xLen, yLen int) grid {
	grid := make(grid, xLen)
	for i := range grid {
		grid[i] = make([]int, yLen)
	}

	return grid
}

type grid [][]int

func (g grid) max() int {
	var max int
	for i := 0; i < len(g); i++ {
		guessed := g.cell(i, i)
		if i == 0 || max < guessed {
			max = guessed
		}
	}

	return max
}

func (g grid) cell(i, j int) int {
	if i < 0 || len(g) <= i {
		return 0
	}
	if len(g) < 1 {
		return 0
	}
	if j < 0 || len(g[0]) <= j {
		return 0
	}

	return g[i][j]
}
