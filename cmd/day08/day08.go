package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type direction int

const (
	up    direction = 0
	down  direction = 1
	right direction = 2
	left  direction = 3
)

// Walk the graph in the specific direction. The `handle` parameter will be fed
// the current values until it returns false or the end of the graph is reached in
// that direction.
func walk(g *util.Grid[int], row, column int, direction direction, handle func(int) bool) {
	var step func(int, int) (int, int)
	switch direction {
	case up:
		step = func(r, c int) (int, int) {
			return r - 1, c
		}
	case down:
		step = func(r, c int) (int, int) {
			return r + 1, c
		}
	case left:
		step = func(r, c int) (int, int) {
			return r, c - 1
		}
	case right:
		step = func(r, c int) (int, int) {
			return r, c + 1
		}
	default:
		panic("invalid direction")
	}

	r, c := step(row, column)
	for {
		if r < 0 || c < 0 || r >= g.Rows() || c >= g.Columns() {
			break
		}

		value := g.Value(r, c)
		if !handle(value) {
			break
		}

		r, c = step(r, c)
	}
}

func isVisible(g *util.Grid[int], row, column int) bool {
	height := g.Value(row, column)
	core := func(direction direction) bool {
		visible := true
		walk(g, row, column, direction, func(value int) bool {
			if value >= height {
				visible = false
				return false
			}

			return true
		})

		return visible
	}

	return core(up) || core(down) || core(left) || core(right)
}

func scenicScore(g *util.Grid[int], row, column int) int {
	height := g.Value(row, column)
	core := func(direction direction) int {
		count := 0
		walk(g, row, column, direction, func(value int) bool {
			count++
			return value < height
		})
		return count
	}

	return core(up) * core(down) * core(left) * core(right)
}

func countVisible(g *util.Grid[int]) int {
	count := 0
	for row := 0; row < g.Rows(); row++ {
		for column := 0; column < g.Columns(); column++ {
			if isVisible(g, row, column) {
				count++
			}
		}
	}

	return count
}

func part1Core(name string) (int, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return -1, err
	}

	g, err := util.ParseGridFromLines(lines)
	if err != nil {
		return -1, err
	}

	count := countVisible(g)
	return count, nil
}

func part1() {
	count, err := part1Core("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Visible trees: %d\n", count)
}

func part2Core(name string) (int, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return -1, err
	}

	g, err := util.ParseGridFromLines(lines)
	if err != nil {
		return -1, err
	}

	max := 0
	for row := 0; row < g.Rows(); row++ {
		for column := 0; column < g.Columns(); column++ {
			score := scenicScore(g, row, column)
			if score > max {
				max = score
			}
		}
	}

	return max, err
}

func part2() {
	max, err := part2Core("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Max scenic score %d\n", max)
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
