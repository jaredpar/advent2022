package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)

	core := func(name string, expected int) {
		count, err := part1Core(name)
		assert.NotError(err)
		assert.EqualInt(expected, count)
	}

	core("example.txt", 21)
	core("input.txt", 1820)
}
