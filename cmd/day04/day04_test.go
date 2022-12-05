package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, value int) {
		pairs, err := parsePairs(f, name)
		assert.NotError(err)
		count := countOverlapsFull(pairs)
		assert.EqualInt(value, count)
	}

	core("example.txt", 2)
	core("input.txt", 547)
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, value int) {
		pairs, err := parsePairs(f, name)
		assert.NotError(err)
		count := countOverlapsAny(pairs)
		assert.EqualInt(value, count)
	}

	core("example.txt", 4)
	core("input.txt", 843)
}
