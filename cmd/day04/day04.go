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

type assignment struct {
	low, high int
}

// Does l completely overlap r
func overlapsFull(l assignment, r assignment) bool {
	return l.high >= r.high && l.low <= r.low
}

// Does l overlap r at all?
func overlapsAny(l assignment, r assignment) bool {
	inRange := func(i int) bool {
		return i >= l.low && i <= l.high
	}

	return inRange(r.low) || inRange(r.high)
}

type pair struct {
	first, second assignment
}

func (p pair) overlapsFull() bool {
	return overlapsFull(p.first, p.second) || overlapsFull(p.second, p.first)
}

func (p pair) overlapsAny() bool {
	return overlapsAny(p.first, p.second) || overlapsAny(p.second, p.first)
}

func newAssignment(low, high int) assignment {
	return assignment{low: low, high: high}
}

func parseAssignment(line string) (a assignment, err error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		err = errors.New("bad format string")
		return
	}

	var low, high int
	low, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	high, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	a = newAssignment(low, high)
	return
}

func parsePairs(f embed.FS, name string) ([]pair, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	pairs := make([]pair, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, errors.New("need two assignments")
		}

		a1, err := parseAssignment(parts[0])
		if err != nil {
			return nil, err
		}

		a2, err := parseAssignment(parts[1])
		if err != nil {
			return nil, err
		}

		pairs[i] = pair{first: a1, second: a2}
	}

	return pairs, nil
}

func countOverlapsFull(pairs []pair) int {
	count := 0
	for _, p := range pairs {
		if p.overlapsFull() {
			count++
		}
	}

	return count
}

func countOverlapsAny(pairs []pair) int {
	count := 0
	for _, p := range pairs {
		if p.overlapsAny() {
			count++
		}
	}

	return count
}

func part1() {
	pairs, err := parsePairs(f, "input.txt")
	if err != nil {
		panic(err)
	}

	count := countOverlapsFull(pairs)
	fmt.Printf("%d\n", count)
}

func part2() {
	pairs, err := parsePairs(f, "input.txt")
	if err != nil {
		panic(err)
	}

	count := countOverlapsAny(pairs)
	fmt.Printf("%d\n", count)
}

func main() {
	p1 := flag.Bool("part 1", false, "run part 1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
