package main

import "fmt"

func main() {}

func dynamic(items map[string]item, limit float64) []string {
	grid := make(grid)
	names, base := names(items), calcCostBase(items)
	var lastName string
	for _, name := range names {
		item := items[name]
		for cost := base; cost <= limit; cost += base {
			var maxCell *cell
			if item.cost <= cost {
				maxCell = &cell{
					val: item.val, names: []string{name},
				}
				if cell := grid.cell(lastName, cost-item.cost); cell != nil {
					maxCell.val += cell.val
					maxCell.names = append(maxCell.names, cell.names...)
				}
			}
			if cell := grid.cell(lastName, cost); cell != nil && (maxCell == nil || maxCell.val < cell.val) {
				maxCell = cell
			}

			grid.setCell(name, cost, maxCell)
		}

		lastName = name
	}

	for _, name := range names {
		fmt.Print(name)
		row := grid[name]
		for cost := base; cost <= limit; cost += base {
			fmt.Print(" ", row[cost])
		}
		fmt.Println()
	}

	return grid.cell(lastName, limit).names
}

func calcCostBase(items map[string]item) float64 {
	base := 1.0
	for _, item := range items {
		dec := item.cost - float64(int(item.cost))
		if dec == 0 {
			continue
		}

		if dec < base {
			base = dec
		}
	}

	return base
}

func names(items map[string]item) []string {
	ns := make([]string, len(items))
	var i int
	for name := range items {
		ns[i] = name
		i++
	}

	return ns
}

type grid map[string]map[float64]*cell

func (g grid) cell(i string, j float64) *cell {
	row := g[i]
	if row == nil {
		return nil
	}

	return row[j]
}

func (g *grid) setCell(i string, j float64, c *cell) {
	if (*g)[i] == nil {
		(*g)[i] = make(map[float64]*cell)
	}

	(*g)[i][j] = c
}

type cell struct {
	val   float64
	names []string
}

type item struct {
	val, cost float64
}
