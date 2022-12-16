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

func part1(name string) int {
	cave := parseCave(name)
	i := 1
	for {
		if cave.dropSand() {
			return i - 1
		}

		i++
	}
}

func main() {
	count := part1("input.txt")
	fmt.Println(count)
}
