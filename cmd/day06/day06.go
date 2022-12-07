package main

import (
	"embed"
	"flag"
	"fmt"
	"sort"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

func getOffsetCore(line string, length int) int {
	length--
	sortBuffer := make([]rune, length)
	buffer := make([]rune, length)
	bufferIndex := 0
	inBuffer := func(r rune) bool {
		for _, e := range buffer {
			if e == r {
				return true
			}
		}

		return false
	}

	bufferHasDupes := func() bool {
		copy(sortBuffer, buffer)
		sort.Slice(sortBuffer, func(i, j int) bool {
			return sortBuffer[i] < sortBuffer[j]
		})

		for i := 0; i+1 < length; i++ {
			if sortBuffer[i] == sortBuffer[i+1] {
				return true
			}
		}

		return false
	}

	for i, r := range line {
		if i >= length && !inBuffer(r) && !bufferHasDupes() {
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

func getPacketOffset(line string) int {
	return getOffsetCore(line, 4)
}

func getMessageOffset(line string) int {
	return getOffsetCore(line, 14)
}

func part1() {
	lines := util.MustReadLines(f, "input.txt")
	start := getPacketOffset(lines[0])
	fmt.Printf("Start is %d\n", start)
}

func part2() {
	lines := util.MustReadLines(f, "input.txt")
	start := getMessageOffset(lines[0])
	fmt.Printf("Start is %d\n", start)
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	if *p1 {
		part1()
	} else {
		part2()
	}
}
