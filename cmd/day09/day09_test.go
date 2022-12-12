package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, expected int) {
		actual, err := part1core(name)
		assert.NotError(err)
		assert.EqualInt(expected, actual)
	}

	core("example.txt", 13)
	core("input.txt", 5695)
}
