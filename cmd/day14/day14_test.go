package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, expected int) {
		assert.EqualInt(expected, part1(name))
	}

	core("example.txt", 24)
	core("input.txt", 1330)
}
