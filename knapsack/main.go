package main

func main() {}

func dynamic(items map[string]*item, limit float64) []string {
	grid, base := make(grid), calcCostBase(items)
	var lastName string
	for _, item := range items {
		for cost := base; cost <= limit; cost += base {
			max := grid.cell(lastName, cost)
			if item.cost <= cost {
				guessed := &cell{
					val: item.val, names: []string{item.name},
				}
				if rest := grid.cell(lastName, cost-item.cost); rest != nil {
					guessed.val += rest.val
					guessed.names = append(guessed.names, rest.names...)
				}

				if max == nil || max.val < guessed.val {
					max = guessed
				}
			}

			grid.setCell(item.name, cost, max)
		}

		lastName = item.name
	}

	return grid.cell(lastName, limit).names
}

func calcCostBase(items map[string]*item) float64 {
	base := 1.0
	for _, item := range items {
		dec := item.cost - float64(int(item.cost))
		if dec != 0 && dec < base {
			base = dec
		}
	}

	return base
}

type grid map[string]map[float64]*cell

func (g grid) cell(name string, cost float64) *cell {
	cs := g[name]
	if cs == nil {
		return nil
	}

	return cs[cost]
}

func (g *grid) setCell(name string, cost float64, c *cell) {
	if (*g)[name] == nil {
		(*g)[name] = make(map[float64]*cell)
	}

	(*g)[name][cost] = c
}

type cell struct {
	val   float64
	names []string
}

type item struct {
	name      string
	val, cost float64
}
