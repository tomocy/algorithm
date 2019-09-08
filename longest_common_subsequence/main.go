package main

func main() {}

func dynamic(vocabs []string, s string) string {
	var longest string
	var max int
	for _, vocab := range vocabs {
		guessed := countCommonSubseqence(vocab, s)
		if max < guessed {
			longest, max = vocab, guessed
		}
	}

	return longest
}

func countCommonSubseqence(a, b string) int {
	grid := newGrid(len(a), len(b))
	for i, p := range a {
		for j, q := range b {
			if i == j && p == q {
				grid[i][j] = grid.cell(i-1, j-1) + 1
				continue
			}

			max := grid.cell(i-1, j)
			if guessed := grid.cell(i, j-1); max < guessed {
				max = guessed
			}

			grid[i][j] = max
		}
	}

	return grid.cell(len(a)-1, len(b)-1)
}

func newGrid(xLen, yLen int) grid {
	grid := make(grid, xLen)
	for i := range grid {
		grid[i] = make([]int, yLen)
	}

	return grid
}

type grid [][]int

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
