package main

func main() {}

func dynamic(items []string, name string) string {
	var longest int
	var longestItem string
	for i, item := range items {
		grid := findCommons(item, name)
		guessed := grid[len(item)-1][len(name)-1]
		if i == 0 || longest < guessed {
			longest = guessed
			longestItem = item
		}
	}

	return longestItem
}

func findCommons(item, name string) grid {
	grid := newGrid(len(item), len(name))
	for i, p := range item {
		for j, q := range name {
			var cell int
			if p == q {
				cell = grid.cell(i-1, j-1) + 1
			} else {
				cell = grid.cell(i-1, j)
				if left := grid.cell(i, j-1); cell < left {
					cell = left
				}
			}

			grid[i][j] = cell
		}
	}

	return grid
}

func newGrid(iLen, jLen int) grid {
	grid := make(grid, iLen)
	for i := range grid {
		grid[i] = make([]int, jLen)
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
