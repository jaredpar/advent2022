package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type point struct {
	row, column int
}

type path struct {
	points []point
}

func (p path) height() int {
	if len(p.points) == 0 {
		return 0
	}

	height := p.points[0].row
	for _, p := range p.points[1:] {
		height = util.Max(height, p.row)
	}

	return height
}

func parsePath(line string) path {
	parts := strings.Split(line, " -> ")
	points := make([]point, 0, len(parts))
	for _, part := range parts {
		segParts := strings.Split(part, ",")
		column := util.StringToInt(segParts[0])
		row := util.StringToInt(segParts[1])
		points = append(points, point{row: row, column: column})
	}

	return path{points: points}
}

func parsePaths(name string) []path {
	paths, err := util.ReadAndParseLines(f, name, func(line string) (path, error) {
		return parsePath(line), nil
	})

	if err != nil {
		panic(err)
	}

	return paths
}

func parseCave(name string) (grid *util.Grid[rune], start int) {
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

	rows := height
	columns := max - min
	grid = util.NewGrid[rune](rows+1, columns+1)
	grid.SetAll('.')
	start = 500 - min
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

	return
}

func stringCave(cave *util.Grid[rune]) string {
	var sb strings.Builder
	for r := 0; r < cave.Rows(); r++ {
		for c := 0; c < cave.Columns(); c++ {
			sb.WriteRune(cave.Value(r, c))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func main() {
	cave, _ := parseCave("example.txt")
	fmt.Println(stringCave(cave))

}
