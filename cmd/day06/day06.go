package main

import (
	"embed"
	"fmt"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func getStartOffset(line string) int {
	buffer := make([]rune, 3)
	bufferIndex := 0
	inBuffer := func(r rune) bool {
		for _, e := range buffer {
			if e == r {
				return true
			}
		}

		return false
	}

	// bubble sort! three elements right now so this is not
	// a big deal but can become so if the number grows
	bufferHasDupes := func() bool {
		for i := 0; i < len(buffer); i++ {
			for j := 0; j < len(buffer); j++ {
				if i != j && buffer[i] == buffer[j] {
					return true
				}
			}
		}

		return false
	}

	for i, r := range line {
		if i >= 3 && !inBuffer(r) && !bufferHasDupes() {
			return i + 1
		}
		buffer[bufferIndex] = r
		bufferIndex++
		if bufferIndex == len(buffer) {
			bufferIndex = 0
		}
	}

	panic("did not find start")
}

func part1() {
	lines := util.MustReadLines(f, "input.txt")
	start := getStartOffset(lines[0])
	fmt.Printf("Start is %d\n", start)
}

func main() {
	part1()

}
