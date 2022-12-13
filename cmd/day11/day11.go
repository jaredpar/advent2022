package main

import (
	"embed"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type monkey struct {
	worryLevels []int
	operation   func(int) int
	testDivisor int
	trueMonkey  int
	falseMonkey int
}

func newMonkey(worryLevels []int, operation func(int) int, testDivisor, trueMonkey, falseMonkey int) *monkey {
	return &monkey{worryLevels: worryLevels, operation: operation, testDivisor: testDivisor, trueMonkey: trueMonkey, falseMonkey: falseMonkey}
}

func stripStart(prefix, input string) (string, error) {
	if util.StartsWithString(input, prefix) {
		return input[len(prefix):], nil
	}

	return "", errors.New("bad line")
}

func parseMonkies(lines []string) ([]*monkey, error) {
	// Parse out a regex where there is a single sub-expression that is
	// converted to a number
	parseNumber := func(pattern, input string) (int, error) {
		r := regexp.MustCompile(pattern)
		indexes := r.FindStringSubmatchIndex(input)
		if len(indexes) != 4 {
			return 0, errors.New("bad line")
		}

		str := input[indexes[2]:indexes[3]]
		return strconv.Atoi(str)
	}

	parseWorryLevels := func(line string) ([]int, error) {
		line, err := stripStart("  Starting items: ", line)
		if err != nil {
			return nil, err
		}

		parts := strings.Split(line, ", ")
		items := make([]int, 0, len(parts))
		for _, part := range parts {
			item, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}

		return items, nil
	}

	parseOperation := func(line string) (operation func(int) int, err error) {
		doubleFunc := func(old int) int {
			return old * old
		}

		getPlusFunc := func(increment int) func(int) int {
			return func(old int) int {
				return old + increment
			}
		}

		getMultiplyFunc := func(increment int) func(int) int {
			return func(old int) int {
				return old * increment
			}
		}

		line, err = stripStart("  Operation: new = old ", line)
		if err != nil {
			return
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			err = errors.New("bad line")
			return
		}

		var number int
		switch parts[0] {
		case "+":
			number, err = strconv.Atoi(parts[1])
			if err != nil {
				return
			}

			operation = getPlusFunc(number)
		case "*":
			if parts[1] == "old" {
				operation = doubleFunc
			} else {
				number, err = strconv.Atoi(parts[1])
				if err != nil {
					return
				}
				operation = getMultiplyFunc(number)
			}
		default:
			err = errors.New("bad operation")
		}

		return
	}

	monkies := make([]*monkey, 0, len(lines)+1/7)
	index := 0
	for index+5 < len(lines) {
		id, err := parseNumber(`Monkey (\d+):`, lines[index])
		if err != nil {
			return nil, err
		}

		if id != len(monkies) {
			return nil, errors.New("bad monkey id")
		}

		worryLevels, err := parseWorryLevels(lines[index+1])
		if err != nil {
			return nil, err
		}

		operation, err := parseOperation(lines[index+2])
		if err != nil {
			return nil, err
		}

		testDivisor, err := parseNumber(`Test: divisible by (\d+)`, lines[index+3])
		if err != nil {
			return nil, err
		}

		trueMonkey, err := parseNumber(`If true: throw to monkey (\d+)`, lines[index+4])
		if err != nil {
			return nil, err
		}

		falseMonkey, err := parseNumber(`If false: throw to monkey (\d+)`, lines[index+5])
		if err != nil {
			return nil, err
		}

		monkey := newMonkey(worryLevels, operation, testDivisor, trueMonkey, falseMonkey)
		monkies = append(monkies, monkey)

		index += 7
	}

	return monkies, nil
}

func readAndParseMonkies(name string) ([]*monkey, error) {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		return nil, err
	}

	return parseMonkies(lines)
}

func main() {
	_, err := readAndParseMonkies("input.txt")
	if err != nil {
		panic(err)
	}
}
