package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS
var debug bool

type direction int

const (
	up    direction = 0
	down  direction = 1
	right direction = 2
	left  direction = 3
)

type move struct {
	direction direction
	count     int
}

func newMove(direction direction, count int) move {
	return move{direction: direction, count: count}
}

type point struct {
	row, column int
}

func newPoint(row, column int) point {
	return point{row: row, column: column}
}

func (p *point) move(d direction) {
	switch d {
	case up:
		p.row--
	case down:
		p.row++
	case left:
		p.column--
	case right:
		p.column++
	default:
		panic("bad direction")
	}
}

func parseMoves(lines []string) ([]move, error) {
	moves := make([]move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, errors.New("bad line")
		}

		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		var d direction
		switch parts[0] {
		case "R":
			d = right
		case "U":
			d = up
		case "L":
			d = left
		case "D":
			d = down
		default:
			return nil, errors.New("bad line")
		}

		moves[i] = newMove(d, count)
	}

	return moves, nil
}

func part1core(name string) (int, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return 0, err
	}

	moves, err := parseMoves(lines)
	if err != nil {
		return 0, err
	}

	head := newPoint(0, 0)
	tail := newPoint(0, 0)
	moveTail := func() {
		if head == tail {
			return
		}

		rowDiff := head.row - tail.row
		columnDiff := head.column - tail.column

		if rowDiff < -1 || rowDiff > 1 {
			if rowDiff < -1 {
				tail.move(up)
			} else if rowDiff > 1 {
				tail.move(down)
			}

			if columnDiff < 0 {
				tail.move(left)
			} else if columnDiff > 0 {
				tail.move(right)
			}
		}

		if columnDiff < -1 || columnDiff > 1 {
			if columnDiff < -1 {
				tail.move(left)
			} else if columnDiff > 1 {
				tail.move(right)
			}

			if rowDiff < 0 {
				tail.move(up)
			} else if rowDiff > 0 {
				tail.move(down)
			}
		}
	}

	printIt := func() {
		top := util.Min(head.row, tail.row)
		bottom := util.Max(head.row, tail.row)
		right := util.Max(head.column, tail.column)
		left := util.Min(head.column, tail.column)

		for r := top; r <= bottom; r++ {
			for c := left; c <= right; c++ {
				current := newPoint(r, c)
				switch current {
				case head:
					fmt.Print("H")
				case tail:
					fmt.Print("T")
				default:
					fmt.Print(".")
				}
			}
			fmt.Println()
		}

		fmt.Println()
	}

	hit := make(map[point]bool)
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			head.move(move.direction)
			moveTail()
			hit[tail] = true

			if debug {
				printIt()
			}
		}
	}

	return len(hit), nil
}

func main() {
	flag.BoolVar(&debug, "debug", false, "visual debug")
	flag.Parse()

	count, err := part1core("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", count)
}
