package main

import (
	"strings"

	"github.com/jaredpar/advent2022/util"
)

type cave struct {
	grid        *util.Grid[rune]
	startColumn int
	hasFloor    bool
}

func (c *cave) String() string {
	var sb strings.Builder
	grid := c.grid
	for r := 0; r < grid.Rows(); r++ {
		for c := 0; c < grid.Columns(); c++ {
			sb.WriteRune(grid.Value(r, c))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (c *cave) isEmpty(row, column int) bool {
	value := c.grid.Value(row, column)
	return value == '.'
}

func (c *cave) dropSand() bool {
	col := c.startColumn
	row := 0
	grid := c.grid

	for {
		if row+1 == grid.Rows() {
			return true
		}
		if c.isEmpty(row+1, col) {
			row++
		} else if col == 0 || col+1 == grid.Columns() {
			if c.hasFloor {
				panic("hit the wall")
			}
			return true
		} else if c.isEmpty(row+1, col-1) {
			row++
			col--
		} else if c.isEmpty(row+1, col+1) {
			row++
			col++
		} else {
			break
		}
	}

	grid.SetValue(row, col, 'o')
	return false
}

func parseCave(name string, hasFloor bool, columnAdjust int) *cave {
	paths := parsePaths(name)
	min := 500
	max := 500
	height := 0
	for _, path := range paths {
		height = util.Max(height, path.height())
		for _, point := range path.points {
			min = util.Min(min, point.column)
			max = util.Max(max, point.column)
		}
	}

	if hasFloor {
		height += 2
	}

	min -= columnAdjust
	max += columnAdjust

	rows := height
	columns := max - min
	grid := util.NewGrid[rune](rows+1, columns+1)
	grid.SetAll('.')
	startColumn := 500 - min
	adjust := func(p point) point {
		p.column -= min
		return p
	}

	// Draw out the cave
	for _, path := range paths {
		if len(path.points) == 0 {
			continue
		}

		current := adjust(path.points[0])
		for _, p := range path.points[1:] {
			p = adjust(p)
			if current.column == p.column {
				start := util.Min(current.row, p.row)
				max := util.Max(current.row, p.row)
				for r := start; r <= max; r++ {
					grid.SetValue(r, current.column, '#')
				}
			} else {
				start := util.Min(current.column, p.column)
				max := util.Max(current.column, p.column)
				for c := start; c <= max; c++ {
					grid.SetValue(current.row, c, '#')
				}
			}

			current = p
		}
	}

	if hasFloor {
		floor := height
		for c := 0; c < grid.Columns(); c++ {
			grid.SetValue(floor, c, '#')
		}
	}

	return &cave{grid: grid, startColumn: startColumn, hasFloor: hasFloor}
}
