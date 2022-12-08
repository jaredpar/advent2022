package main

import (
	"embed"
	"fmt"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func isVisible(g *util.Grid[int], row, column int) bool {
	if row == 0 || column == 0 {
		return true
	}

	if row+1 == g.Rows() || column+1 == g.Columns() {
		return true
	}

	height := g.Value(row, column)
	core := func(step func(int, int) (int, int)) bool {
		r, c := step(row, column)
		for {
			if r < 0 || c < 0 || r >= g.Rows() || c >= g.Columns() {
				return true
			}

			if g.Value(r, c) >= height {
				return false
			}

			r, c = step(r, c)
		}
	}

	stepLeft := func(r, c int) (int, int) {
		return r, c - 1
	}
	stepRight := func(r, c int) (int, int) {
		return r, c + 1
	}
	stepUp := func(r, c int) (int, int) {
		return r - 1, c
	}
	stepDown := func(r, c int) (int, int) {
		return r + 1, c
	}
	return core(stepLeft) || core(stepRight) || core(stepUp) || core(stepDown)
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

func main() {
	part1()
}
