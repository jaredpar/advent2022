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

type assignment struct {
	low, high int
}

// Does l completely overlap r
func overlaps(l assignment, r assignment) bool {
	return l.high >= r.high && l.low <= r.low
}

type pair struct {
	first, second assignment
}

func (p pair) overlaps() bool {
	return overlaps(p.first, p.second) || overlaps(p.second, p.first)
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

func countOverlaps(pairs []pair) int {
	count := 0
	for _, p := range pairs {
		if p.overlaps() {
			count++
		}
	}

	return count
}

func main() {
	pairs, err := parsePairs(f, "input.txt")
	if err != nil {
		panic(err)
	}

	count := countOverlaps(pairs)
	fmt.Printf("%d\n", count)
}
