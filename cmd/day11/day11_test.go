package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	assert.EqualInt(10605, part1("example.txt"))
	assert.EqualInt(54036, part1("input.txt"))
}
