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

func run(instructions []instruction, callback func(int, int)) int {
	var currentIndex, remaining int
	var current *instruction
	changeCurrent := func(index int) {
		currentIndex = index
		current = &instructions[currentIndex]
		remaining = current.cycles()
	}
	changeCurrent(0)

	register := 1
	sum := 0
	cycle := 1

	for {
		callback(cycle, register)

		remaining--
		if remaining == 0 {
			if instructions[currentIndex].kind == addx {
				register += instructions[currentIndex].count
			}

			if currentIndex+1 == len(instructions) {
				break
			}

			changeCurrent(currentIndex + 1)
		}
		cycle++
	}

	return sum
}

func getSum(instructions []instruction) int {
	nextCycleCheck := 20
	sum := 0

	run(instructions, func(cycle, register int) {
		if cycle == nextCycleCheck {
			sum += register * cycle
			nextCycleCheck += 40
		}
	})

	return sum
}

func draw(instructions []instruction) string {
	var sb strings.Builder
	run(instructions, func(cycle, register int) {
		column := (cycle % 40) - 1
		if (register-1) <= column && (register+1) >= column {
			sb.WriteRune('#')
		} else {
			sb.WriteRune('.')
		}

		if cycle%40 == 0 {
			sb.WriteRune('\n')
		}
	})
	return sb.String()
}

func part1Core(name string) (int, error) {
	instructions, err := parseInstructions(name)
	if err != nil {
		return 0, err
	}

	return getSum(instructions), nil
}

func part2Core(name string) (string, error) {
	instructions, err := parseInstructions(name)
	if err != nil {
		return "", err
	}

	return draw(instructions), nil
}

func main() {
	p1 := flag.Bool("part1", false, "run part 1")
	if *p1 {
		sum, err := part1Core("input.txt")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d\n", sum)
	} else {
		screen, err := part2Core("input.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println(screen)
	}
}
