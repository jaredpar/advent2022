package main

import (
	"embed"
	"fmt"
	"flag"
	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS
var debug bool

func getSumOfMax(count int) int {
	maxes := make([]int, count)
	current := 0
	lines := util.MustReadLines(f, "input1.txt")
	for _, line := range lines {
		if util.IsWhitespace(line) {
			for i, max := range maxes {
				if current > max {
					maxes[i] = current
					break
				}
			}

			current = 0
			continue
		}

		value := util.StringToInt(line)
		current = current + value
	}

	total := 0
	for _, max := range maxes {
		total += max
	}

	return total
}

func part1() {
	value := getSumOfMax(1)
	fmt.Printf("Max calories %d\n", value)
}

func part2() {
	value := getSumOfMax(3)
	fmt.Printf("Max calories %d\n", value)
}

func main() {
	p1 := flag.Bool("part1", false, "run part1")
	flag.Parse()
	if *p1 {
		part1()
	} else {
		part2()
	}
}
