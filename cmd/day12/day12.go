package main

import (
	"embed"
	"fmt"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type point struct {
	row, column int
}

func newPoint(row, column int) point {
	return point{row: row, column: column}
}

func parseInput(name string) (grid *util.Grid[int], start point, end point) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		panic(err)
	}

	grid, err = util.ParseGridFromLines(lines, func(row, col int, r rune) (int, error) {
		switch r {
		case 'S':
			start = newPoint(row, col)
			return 0, nil
		case 'E':
			end = newPoint(row, col)
			return 25, nil
		default:
			return int(r - 'a'), nil
		}
	})

	if err != nil {
		panic(err)
	}

	return
}

type visitInfo struct {
	// Have we visited this node to look at it's edges
	visited bool

	// what is the minimum distance to this node
	distance int
}

func newVisitInfo(distance int) *visitInfo {
	return &visitInfo{visited: false, distance: distance}
}

func part1(name string) int {
	grid, start, end := parseInput(name)
	visitMap := make(map[point]*visitInfo)
	visitMap[start] = newVisitInfo(0)

	getNextVisit := func() (next point, ok bool) {
		any := false
		var minDistance int

		for key, value := range visitMap {
			if !value.visited {
				if !any || value.distance < minDistance {
					minDistance = value.distance
					next = key
					any = true
				}
			}
		}

		ok = any
		return
	}

	// Can from visit to? The to point is possibly outside the bounds
	// of the graph while from will always be a legal point
	canVisit := func(from, to point) bool {
		if to.row < 0 || to.column < 0 || to.row >= grid.Rows() || to.column >= grid.Columns() {
			return false
		}

		fromHeight := grid.Value(from.row, from.column)
		toHeight := grid.Value(to.row, to.column)
		return fromHeight+1 >= toHeight
	}

	queueEdges := func(from point, fromDistance int) {
		queueOne := func(p point) {
			if !canVisit(from, p) {
				return
			}

			// Already visited
			if _, ok := visitMap[p]; ok {
				return
			}

			visitMap[p] = newVisitInfo(fromDistance + 1)
		}

		queueOne(newPoint(from.row-1, from.column))
		queueOne(newPoint(from.row+1, from.column))
		queueOne(newPoint(from.row, from.column-1))
		queueOne(newPoint(from.row, from.column+1))
	}

	for {
		current, ok := getNextVisit()
		if !ok {
			break
		}

		currentInfo, ok := visitMap[current]
		if !ok {
			panic("node should be visited")
		}

		currentInfo.visited = true
		queueEdges(current, currentInfo.distance)
	}

	endDist, ok := visitMap[end]
	if !ok {
		panic("never reached end")
	}

	return endDist.distance
}

func main() {
	steps := part1("input.txt")
	fmt.Println(steps)
}
