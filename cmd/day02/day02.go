package main

import (
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

type move int

const (
	rock     move = 1
	paper    move = 2
	scissors move = 3
)

type round struct {
	opponent move
	your     move
}

func (r round) isTie() bool {
	return r.opponent == r.your
}

func (r round) isWin() bool {
	switch r.your {
	case rock:
		return r.opponent == scissors
	case paper:
		return r.opponent == rock
	case scissors:
		return r.opponent == paper
	default:
		panic("invalid move")
	}
}

func (r round) isLoss() bool {
	return !r.isWin() && !r.isTie()
}

// Get the score of the round based on the formula in the problem
func (r round) getScore() int {
	shape := int(r.your)
	if r.isTie() {
		return 3 + shape
	} else if r.isWin() {
		return 6 + shape
	} else {
		return shape
	}
}

func newRound(opponent, your move) round {
	return round{opponent: opponent, your: your}
}

func parseInput(f embed.FS, name string) ([]round, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	rounds := make([]round, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, errors.New("invalid line")
		}

		var opponent, your move
		switch parts[0] {
		case "A":
			opponent = rock
		case "B":
			opponent = paper
		case "C":
			opponent = scissors
		default:
			return nil, errors.New("bad opponent move")
		}

		switch parts[1] {
		case "X":
			your = rock
		case "Y":
			your = paper
		case "Z":
			your = scissors
		default:
			return nil, errors.New("bad your move")
		}

		rounds[i] = newRound(opponent, your)
	}

	return rounds, nil
}

func getTotalScore(rounds []round) int {
	total := 0
	for _, round := range rounds {
		total += round.getScore()
	}

	return total
}

//go:embed *.txt
var f embed.FS

func part1() {
	rounds, err := parseInput(f, "input.txt")
	if err != nil {
		panic("could not parse input")
	}

	fmt.Printf("Total: %d\n", getTotalScore(rounds))
}

func main() {
	part1()
}
