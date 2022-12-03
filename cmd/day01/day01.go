package main

import (
	"embed"
	"fmt"
	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func part1() {
	max := 0
	current := 0
	lines := util.MustReadLines(f, "input1.txt")
	for _, line := range lines {
		if util.IsWhitespace(line) {
			fmt.Printf("Current calories %d\n", current)
			if current > max {
				max = current
			}

			current = 0
			continue
		}

		value := util.StringToInt(line)
		current = current + value
	}

	fmt.Printf("Max calories is %d\n", max)
}

func main() {
	part1()

}
