package main

import "github.com/jaredpar/advent2022/util"

type packet struct {
	value any
}

func newPacketSingle(s int) packet {
	return packet{value: s}
}

func newPacketList(values []packet) packet {
	return packet{value: values}
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
