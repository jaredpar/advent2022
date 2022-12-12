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

type rope struct {
	points []point
}

func newRope(knots int) *rope {
	if knots < 2 {
		panic("must be at least 2 knots")
	}

	p := make([]point, knots)
	return &rope{points: p}
}

func (r *rope) move(direction direction) {
	moveNext := func(head, tail *point) {
		if *head == *tail {
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
		} else if columnDiff < -1 || columnDiff > 1 {
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

	r.points[0].move(direction)
	for i := 1; i < len(r.points); i++ {
		moveNext(&r.points[i-1], &r.points[i])
	}
}

func (r *rope) moveCount(d direction, count int) {
	for i := 0; i < count; i++ {
		r.move(d)
	}
}

func (r *rope) head() *point {
	return &r.points[0]
}

func (r *rope) tail() *point {
	return &r.points[len(r.points)-1]
}

func (r *rope) String() string {
	var sb strings.Builder
	top := r.points[0].row
	bottom := top
	left := r.points[0].column
	right := left

	for i := 1; i < len(r.points); i++ {
		current := r.points[i]
		top = util.Min(top, current.row)
		bottom = util.Max(bottom, current.row)
		left = util.Min(left, current.column)
		right = util.Max(right, current.column)
	}

	for row := top; row <= bottom; row++ {
		for column := left; column <= right; column++ {
			position := newPoint(row, column)
			switch position {
			case *r.head():
				sb.WriteRune('H')
			case *r.tail():
				sb.WriteRune('T')
			default:
				any := false
				for i := 1; i+1 < len(r.points); i++ {
					if r.points[i] == position {
						if i < 10 {
							fmt.Fprintf(&sb, "%d", i)
						} else {
							sb.WriteRune('*')
						}

						any = true
						break
					}
				}

				if !any {
					sb.WriteRune('.')
				}
			}
		}

		sb.WriteRune('\n')
	}

	return sb.String()
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

func runCore(name string, knots int) (int, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return 0, err
	}

	moves, err := parseMoves(lines)
	if err != nil {
		return 0, err
	}

	rope := newRope(knots)
	hit := make(map[point]bool)
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			rope.move(move.direction)
			hit[*rope.tail()] = true
			if debug {
				fmt.Println(rope)
			}
		}
	}

	return len(hit), nil
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	flag.BoolVar(&debug, "debug", false, "visual debug")
	flag.Parse()

	var knots int
	if *p1 {
		knots = 2
	} else {
		knots = 10
	}

	count, err := runCore("input.txt", knots)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", count)
}
