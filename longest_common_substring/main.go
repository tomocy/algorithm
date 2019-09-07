package main

func main() {}

func dynamic(items []string, name string) string {
	var longestCell int
	var longestItem string
	for i, item := range items {
		grid := findCommons(item, name)
		cell := grid.longestCell()
		if i == 0 || longestCell < cell {
			longestCell = cell
			longestItem = item
		}
	}

	return longestItem
}

func findCommons(item, name string) grid {
	grid := newGrid(len(item), len(name))
	for i := 0; i < len(item) && i < len(name); i++ {
		var cell int
		if item[i] == name[i] {
			cell = grid.cell(i-1, i-1) + 1
		}

		grid[i][i] = cell
	}

	return grid
}

func newGrid(lenI, lenJ int) grid {
	grid := make([][]int, lenI)
	for i := range grid {
		grid[i] = make([]int, lenJ)
	}

	return grid
}

type grid [][]int

func (g grid) cell(i, j int) int {
	if i < 0 && len(g) <= i {
		return 0
	}
	if len(g) < 1 {
		return 0
	}
	if j < 0 && len(g[0]) <= j {
		return 0
	}

	return g[i][j]
}

func (g grid) longestCell() int {
	var longest, i int
	for _, row := range g {
		for _, cell := range row {
			if i == 0 || longest < cell {
				longest = cell
			}

			i++
		}
	}

	return longest
}
