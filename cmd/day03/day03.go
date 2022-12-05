package main

import (
	"embed"
	"errors"
	"flag"
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

func sumBadges(sacks []*sack) int {
	getMap := func(s *sack) map[rune]bool {
		m := make(map[rune]bool)
		for _, r := range s.first {
			m[r] = true
		}
		for _, r := range s.second {
			m[r] = true
		}

		return m
	}

	getBadge := func(sacks []*sack) rune {
		m1 := getMap(sacks[0])
		m2 := getMap(sacks[1])
		inBoth := func(r rune) bool {
			_, p1 := m1[r]
			_, p2 := m2[r]
			return p1 && p2
		}

		sack := sacks[2]
		for _, r := range sack.first {
			if inBoth(r) {
				return r
			}
		}
		for _, r := range sack.second {
			if inBoth(r) {
				return r
			}
		}

		panic("no badge")
	}

	index := 0
	sum := 0
	for index < len(sacks) {
		r := getBadge(sacks[index : index+3])
		sum += priority(r)
		index += 3
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

func part2() {
	sacks, err := parseSacks(f, "input.txt")
	if err != nil {
		panic("bad example")
	}

	sum := sumBadges(sacks)
	fmt.Printf("Sum is %d\n", sum)
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
