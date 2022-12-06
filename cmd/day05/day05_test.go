package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	ship, moves, err := parseInput(f, "input.txt")
	assert.NotError(err)
	ship.runMoves(moves)
	assert.EqualString("ZBDRNPMVH", ship.message())
}
