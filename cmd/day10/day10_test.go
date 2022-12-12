package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)

	core := func(name string, expected int) {
		sum, err := part1Core(name)
		assert.NotError(err)
		assert.EqualInt(expected, sum)
	}

	core("example.txt", 13140)
	core("input.txt", 15360)
}
