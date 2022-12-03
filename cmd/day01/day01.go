package main

import (
	"embed"
	"fmt"
	"flag"
	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func getSumOfMax(count int) {
	maxes := make([]int, count)
	current := 0
	lines := util.MustReadLines(f, "input1.txt")
	for _, line := range lines {
		if util.IsWhitespace(line) {
			fmt.Printf("Current calories %d\n", current)
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
		fmt.Printf("Individual max is %d\n", max)
		total += max
	}

	fmt.Printf("Max calories is %d\n", total)
}

func part1() {
	getSumOfMax(1)
}

func part2() {
	getSumOfMax(3)
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
