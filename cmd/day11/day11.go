package main

import (
	"embed"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

//go:embed *.txt
var f embed.FS

type monkey struct {
	items       []int
	operation   func(int) int
	testDivisor int
	trueMonkey  int
	falseMonkey int
	inspected   int
}

func newMonkey(items []int, operation func(int) int, testDivisor, trueMonkey, falseMonkey int) *monkey {
	return &monkey{items: items, operation: operation, testDivisor: testDivisor, trueMonkey: trueMonkey, falseMonkey: falseMonkey}
}

type troop []*monkey

func (t troop) turn(id int) {
	current := t[id]
	items := current.items
	current.items = make([]int, 0)
	for _, worry := range items {
		current.inspected++
		worry = current.operation(worry) / 3

		var target *monkey
		if worry%current.testDivisor == 0 {
			target = t[current.trueMonkey]
		} else {
			target = t[current.falseMonkey]
		}

		target.items = append(target.items, worry)
	}
}

func (t troop) round() {
	for i, _ := range t {
		t.turn(i)
	}
}

func (t troop) String() string {
	var sb strings.Builder
	for i, m := range t {
		fmt.Fprintf(&sb, "Monkey %d: ", i)
		first := true
		for _, level := range m.items {
			if !first {
				sb.WriteString(", ")
			}
			fmt.Fprintf(&sb, "%d", level)
			first = false
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func stripStart(prefix, input string) string {
	if util.StartsWithString(input, prefix) {
		return input[len(prefix):]
	}

	panic("bad line")
}

func parseMonkeys(lines []string) []*monkey {
	// Parse out a regex where there is a single sub-expression that is
	// converted to a number
	parseNumber := func(pattern, input string) int {
		r := regexp.MustCompile(pattern)
		indexes := r.FindStringSubmatchIndex(input)
		if len(indexes) != 4 {
			panic("bad line")
		}

		str := input[indexes[2]:indexes[3]]
		return util.StringToInt(str)
	}

	parseItems := func(line string) []int {
		line = stripStart("  Starting items: ", line)
		parts := strings.Split(line, ", ")
		items := make([]int, 0, len(parts))
		for _, part := range parts {
			item := util.StringToInt(part)
			items = append(items, item)
		}

		return items
	}

	parseOperation := func(line string) func(int) int {
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

		line = stripStart("  Operation: new = old ", line)
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			panic("bad line")
		}

		var number int
		switch parts[0] {
		case "+":
			number := util.StringToInt(parts[1])
			return getPlusFunc(number)
		case "*":
			if parts[1] == "old" {
				return doubleFunc
			} else {
				number = util.StringToInt(parts[1])
				return getMultiplyFunc(number)
			}
		default:
			panic("bad operator")
		}
	}

	monkeys := make([]*monkey, 0, len(lines)+1/7)
	index := 0
	for index+5 < len(lines) {
		id := parseNumber(`Monkey (\d+):`, lines[index])
		if id != len(monkeys) {
			panic("bad monkey id")
		}

		items := parseItems(lines[index+1])
		operation := parseOperation(lines[index+2])
		testDivisor := parseNumber(`Test: divisible by (\d+)`, lines[index+3])
		trueMonkey := parseNumber(`If true: throw to monkey (\d+)`, lines[index+4])
		falseMonkey := parseNumber(`If false: throw to monkey (\d+)`, lines[index+5])
		monkey := newMonkey(items, operation, testDivisor, trueMonkey, falseMonkey)
		monkeys = append(monkeys, monkey)

		index += 7
	}

	return monkeys
}

func readAndParseMonkeys(name string) []*monkey {
	lines, err := util.ReadLines(f, name)
	if err != nil {
		panic(err)
	}

	return parseMonkeys(lines)
}

func readAndParseTroop(name string) troop {
	return readAndParseMonkeys(name)
}

func part1(name string) int {
	troop := readAndParseTroop(name)

	for r := 0; r < 20; r++ {
		troop.round()
		fmt.Println(troop)
	}

	counts := util.Project(troop, func(m *monkey) int {
		return m.inspected
	})
	sort.Ints(counts)
	counts = counts[len(counts)-2:]
	return counts[0] * counts[1]
}

func main() {
	count := part1("input.txt")
	fmt.Println(count)
}
