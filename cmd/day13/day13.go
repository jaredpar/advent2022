package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredpar/advent2022/util"
)

type packet struct {
	value any
}

func newPacketSingle(s int) packet {
	return packet{value: s}
}

func newPacketList(values []packet) packet {
	return packet{value: values}
}

func (p packet) isList() bool {
	_, ok := p.value.([]packet)
	return ok
}

func (p packet) isSingle() bool {
	_, ok := p.value.(int)
	return ok
}

func compareSlice(left, right []packet) int {
	count := util.Max(len(left), len(right))
	i := 0
	for i < count {
		if len(left) != len(right) {
			if i == len(left) {
				return -1
			}

			if i == len(right) {
				return 1
			}
		}

		comp := left[i].compare(right[i])
		if comp != 0 {
			return comp
		}

		i++
	}

	return 0
}

func (p packet) compare(other packet) int {
	switch l := p.value.(type) {
	case int:
		switch r := other.value.(type) {
		case int:
			return l - r
		case []packet:
			return compareSlice([]packet{p}, r)
		}
	case []packet:
		switch r := other.value.(type) {
		case int:
			return compareSlice(l, []packet{other})
		case []packet:
			return compareSlice(l, r)
		}
	}
	panic("bad types")
}

func stringCore(sb *strings.Builder, p packet) {
	switch l := p.value.(type) {
	case int:
		fmt.Fprintf(sb, "%d", l)
	case []packet:
		sb.WriteRune('[')
		first := true
		for _, c := range l {
			if !first {
				sb.WriteString(", ")
			}

			stringCore(sb, c)
			first = false
		}
		sb.WriteRune(']')
	}
}

func (p packet) String() string {
	var sb strings.Builder
	stringCore(&sb, p)
	return sb.String()
}

func parsePacket(line string) packet {
	parseOne := func(runes []rune) (packet, []rune) {
		end := 1
		for end < len(runes) {
			if runes[end] == ',' || runes[end] == ']' {
				break
			}

			if end+1 == len(runes) {
				break
			}

			end++
		}

		var d int
		var e error
		if end == 1 {
			d, e = util.RuneToInt(runes[0])
		} else {
			str := string(runes[0:end])
			d, e = strconv.Atoi(str)
		}

		if e != nil {
			panic("bad rune")
		}
		return newPacketSingle(d), runes[end:]
	}

	var parseList func([]rune) (packet, []rune)
	parseList = func(runes []rune) (packet, []rune) {
		runes = runes[1:]
		packets := make([]packet, 0)

		for len(runes) > 0 {
			switch runes[0] {
			case '[':
				p, rest := parseList(runes)
				runes = rest
				packets = append(packets, p)
			case ']':
				runes = runes[1:]
				break
			case ',':
				runes = runes[2:]
			default:
				p, rest := parseOne(runes)
				runes = rest
				packets = append(packets, p)
			}
		}

		return newPacketList(packets), runes
	}

	p, rest := parseList([]rune(line))
	if len(rest) != 0 {
		panic("extra items")
	}

	return p
}
