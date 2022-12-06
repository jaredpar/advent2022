package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	ship, moves, err := parseInput(f, "input.txt")
	assert.NotError(err)
	ship.runMoves(moves /* keepOrder */, false)
	assert.EqualString("ZBDRNPMVH", ship.message())
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	ship, moves, err := parseInput(f, "input.txt")
	assert.NotError(err)
	ship.runMoves(moves /* keepOrder */, true)
	assert.EqualString("WDLPFNNNB", ship.message())
}
