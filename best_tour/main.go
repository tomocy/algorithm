package main

func main() {}

func dynamic(trips map[string]*trip, limit float64) []string {
	grid, base := make(grid), calcCostBase(trips)
	var lastName string
	for _, trip := range trips {
		for cost := base; cost <= limit; cost += base {
			max := grid.cell(lastName, cost)
			if trip.cost <= cost {
				guessed := &cell{
					val: trip.val, names: []string{trip.name},
				}
				if rest := grid.cell(lastName, cost-trip.cost); rest != nil {
					guessed.val += rest.val
					guessed.names = append(guessed.names, rest.names...)
				}

				if max == nil || max.val < guessed.val {
					max = guessed
				}
			}

			grid.setCell(trip.name, cost, max)
		}

		lastName = trip.name
	}

	return grid[lastName][limit].names
}

func calcCostBase(trips map[string]*trip) float64 {
	base := 1.0
	for _, trip := range trips {
		dec := trip.cost - float64(int(trip.cost))
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

type trip struct {
	name      string
	val, cost float64
}
