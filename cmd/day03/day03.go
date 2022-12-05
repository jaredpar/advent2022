package main

import (
	"embed"
	"errors"
	"fmt"
	"unicode"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func priority(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - int('a') + 1
	} else {
		return (int(r) - int('A')) + 27
	}
}

type sack struct {
	first, second string
}

func (s sack) shared() []rune {
	m := make(map[rune]bool)
	for _, r := range s.first {
		m[r] = true
	}

	var sh []rune
	for _, r := range s.second {
		if _, present := m[r]; present {
			delete(m, r)
			sh = append(sh, r)
		}
	}

	return sh
}

func parseSack(line string) (*sack, error) {
	length := len(line)
	if length%2 != 0 {
		return nil, errors.New("line not even number of characters")
	}

	mid := length / 2
	return &sack{first: line[0:mid], second: line[mid:]}, nil
}

func parseSacks(f embed.FS, name string) ([]*sack, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	sacks := make([]*sack, len(lines))
	for i, line := range lines {
		sack, err := parseSack(line)
		if err != nil {
			return nil, err
		}

		sacks[i] = sack
	}

	return sacks, nil
}

func sumShared(sacks []*sack) int {
	sum := 0
	for _, sack := range sacks {
		for _, r := range sack.shared() {
			sum += priority(r)
		}
	}

	return sum
}

func part1() {
	sacks, err := parseSacks(f, "input.txt")
	if err != nil {
		panic("bad example")
	}

	sum := sumShared(sacks)
	fmt.Printf("Sum is %d\n", sum)
}

func main() {
	part1()
}
