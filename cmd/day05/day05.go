package main

import (
	"embed"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type move struct {
	count, from, to int
}

func (m move) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.count, m.from, m.to)
}

func newMove(count, from, to int) move {
	return move{count: count, from: from, to: to}
}

type ship struct {
	stacks []util.Stack[rune]
}

func (s *ship) runMove(m move, keepOrder bool) {
	fs := &s.stacks[m.from-1]
	ts := &s.stacks[m.to-1]
	if keepOrder {
		temp := util.NewStack[rune]()
		c := 0
		for c < m.count {
			temp.Push(fs.Pop())
			c++
		}

		for temp.Length() > 0 {
			ts.Push(temp.Pop())
		}
	} else {
		c := 0
		for c < m.count {
			ts.Push(fs.Pop())
			c++
		}
	}
}

func (s *ship) runMoves(moves []move, keepOrder bool) {
	for _, m := range moves {
		s.runMove(m, keepOrder)
	}
}

func (s *ship) message() string {
	var sb strings.Builder
	for _, stack := range s.stacks {
		sb.WriteRune(stack.Peek())
	}
	return sb.String()
}

func parseShip(lines []string) ship {

	// First calculate the length of the crate system.
	length := (len(lines[0]) + 1) / 4
	stacks := make([]util.Stack[rune], length)

	// Parse out every row
	for _, line := range lines {
		lineIndex := 0
		index := 0
		for lineIndex < len(line) {
			item := line[lineIndex : lineIndex+3]
			runes := []rune(item)
			if runes[0] == '[' {
				stacks[index].Push(runes[1])
			}

			lineIndex += 4
			index++
		}
	}

	for _, stack := range stacks {
		stack.Reverse()
	}

	return ship{stacks: stacks}
}

func parseMoves(lines []string) ([]move, error) {
	moves := make([]move, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 6 ||
			parts[0] != "move" ||
			parts[2] != "from" ||
			parts[4] != "to" {
			return nil, errors.New("bad line")
		}

		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		from, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(parts[5])
		if err != nil {
			return nil, err
		}

		move := newMove(count, from, to)
		moves = append(moves, move)
	}

	return moves, nil
}

func parseInput(f embed.FS, name string) (*ship, []move, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, nil, err
	}

	blankIndex := 0
	for i, line := range lines {
		if util.IsWhitespace(line) {
			blankIndex = i
			break
		}
	}

	ship := parseShip(lines[0 : blankIndex-1])

	moves, err := parseMoves(lines[blankIndex+1:])
	if err != nil {
		return nil, nil, err
	}

	return &ship, moves, nil
}

func main() {
	ship, moves, err := parseInput(f, "input.txt")
	if err != nil {
		panic(err)
	}

	ship.runMoves(moves /* keepOrder */, true)
	fmt.Println(ship.message())
}
