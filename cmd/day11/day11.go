package main

import (
	"embed"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type monkey struct {
	worryLevels []int
	worryFunc   func(int) int
	testDivisor int
	trueMonkey  int
	falseMonkey int
}

func newMonkey(worryLevels []int, worryFunc func(int) int, testDivisor, trueMonkey, falseMonkey int) *monkey {
	return &monkey{worryLevels: worryLevels, worryFunc: worryFunc, testDivisor: testDivisor, trueMonkey: trueMonkey, falseMonkey: falseMonkey}
}

func stripStart(prefix string, input string) (string, error) {
	pattern := fmt.Sprintf("^%s(.*)", prefix)
	r := regexp.MustCompile(pattern)
	indexes := r.FindStringSubmatchIndex(input)
	if len(indexes) != 4 {
		return "", errors.New("bad line")
	}

	return input[indexes[2]:indexes[3]], nil
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

		var worryFunc func(int) int
		// TODO: parse this

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

		monkey := newMonkey(worryLevels, worryFunc, testDivisor, trueMonkey, falseMonkey)
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
