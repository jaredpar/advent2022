package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestShared(t *testing.T) {
	assert := testUtil.NewAssert(t)
	sack, err := parseSack("abaa")
	assert.NotError(err)
	assert.EqualInt(1, len(sack.shared()))
}

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, value int) {
		sacks, err := parseSacks(f, name)
		assert.NotError(err)
		sum := sumShared(sacks)
		assert.EqualInt(value, sum)
	}

	core("example.txt", 157)
	core("input.txt", 7903)
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, value int) {
		sacks, err := parseSacks(f, name)
		assert.NotError(err)
		sum := sumBadges(sacks)
		assert.EqualInt(value, sum)
	}

	core("example.txt", 70)
	core("input.txt", 2548)
}
