package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestSample(t *testing.T) {
	assert := testUtil.NewAssert(t)
	rounds, err := parseInput(f, "testInput.txt")
	assert.NotError(err)
	assert.EqualInt(15, getTotalScore(rounds))
}

func TestSample2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	rounds, err := parseInput(f, "testInput.txt")
	decode(rounds)
	assert.NotError(err)
	assert.EqualInt(12, getTotalScore(rounds))
}

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	rounds, err := parseInput(f, "input.txt")
	assert.NotError(err)
	assert.EqualInt(12586, getTotalScore(rounds))
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	rounds, err := parseInput(f, "input.txt")
	assert.NotError(err)
	decode(rounds)
	assert.EqualInt(13193, getTotalScore(rounds))
}
