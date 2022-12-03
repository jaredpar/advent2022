package main

import (
	"embed"
	"errors"
	"flag"
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

func (m move) getWinsAgainst() move {
	switch m {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	default:
		panic("not a valid move")
	}
}

func (m move) getLosesAgainst() move {
	switch m {
	case rock:
		return paper
	case paper:
		return scissors
	case scissors:
		return rock
	default:
		panic("not a valid move")
	}
}

type round struct {
	opponent move
	your     move
}

func (r round) isTie() bool {
	return r.opponent == r.your
}

func (r round) isWin() bool {
	return r.your.getWinsAgainst() == r.opponent
}

func (r round) isLoss() bool {
	return r.your.getLosesAgainst() == r.opponent
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

// Decode according to the pattern
// X == rock == lose
// Y == paper == tie
// Z == scissor == win
func (r *round) decode() {
	switch r.your {
	case rock:
		r.your = r.opponent.getWinsAgainst()
	case paper:
		r.your = r.opponent
	case scissors:
		r.your = r.opponent.getLosesAgainst()
	default:
		panic("not a move")
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

func decode(rounds []round) {
	for i := range rounds {
		rounds[i].decode()
	}
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

func part2() {
	rounds, err := parseInput(f, "input.txt")
	if err != nil {
		panic("could not parse input")
	}
	decode(rounds)
	fmt.Printf("Total: %d\n", getTotalScore(rounds))
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
