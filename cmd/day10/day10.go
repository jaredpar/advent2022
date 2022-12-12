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

type kind int

const (
	noop kind = 0
	addx kind = 1
)

type instruction struct {
	kind  kind
	count int
}

func (i instruction) cycles() int {
	switch i.kind {
	case noop:
		return 1
	case addx:
		return 2
	default:
		panic("bad kind")
	}
}

func newInstruction(kind kind, count int) instruction {
	return instruction{kind: kind, count: count}
}

func parseInstructions(name string) ([]instruction, error) {
	return util.ReadAndParseLines(f, name, func(line string) (inst instruction, err error) {
		parts := strings.Split(line, " ")
		switch len(parts) {
		case 1:
			if parts[0] == "noop" {
				inst = newInstruction(noop, 0)
				return
			}
		case 2:
			if parts[0] == "addx" {
				var count int
				count, err = strconv.Atoi(parts[1])
				if err == nil {
					inst = newInstruction(addx, count)
					return
				}
			}
		}

		if err == nil {
			err = errors.New("bad line")
		}

		return
	})
}

func getSum(instructions []instruction) int {
	var currentIndex, remaining int
	var current *instruction
	changeCurrent := func(index int) {
		currentIndex = index
		current = &instructions[currentIndex]
		remaining = current.cycles()
	}
	changeCurrent(0)

	register := 1
	nextCycleCheck := 20
	sum := 0

	for cycle := 1; cycle <= 220; cycle++ {
		if cycle == nextCycleCheck {
			sum += register * cycle
			nextCycleCheck += 40
		}

		remaining--
		if remaining == 0 {
			if instructions[currentIndex].kind == addx {
				register += instructions[currentIndex].count
			}

			changeCurrent(currentIndex + 1)
		}
	}

	return sum
}

func part1Core(name string) (int, error) {
	instructions, err := parseInstructions(name)
	if err != nil {
		return 0, err
	}

	return getSum(instructions), nil
}

func main() {
	sum, err := part1Core("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", sum)
}
