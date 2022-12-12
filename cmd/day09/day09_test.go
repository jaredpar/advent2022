package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, expected int) {
		actual, err := runCore(name, 2)
		assert.NotError(err)
		assert.EqualInt(expected, actual)
	}

	core("example.txt", 13)
	core("input.txt", 5695)
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	core := func(name string, expected int) {
		actual, err := runCore(name, 10)
		assert.NotError(err)
		assert.EqualInt(expected, actual)
	}

	core("example.txt", 1)
	core("example2.txt", 36)
	core("input.txt", 2434)
}

func TestMove(t *testing.T) {
	assert := testUtil.NewAssert(t)
	rope := newRope(9)
	moveCount := func(d direction, count int) {
		for i := 0; i < count; i++ {
			rope.move(d)
			t.Log(rope)
		}
	}

	moveCount(right, 4)
	assert.EqualString("T321H\n", rope.String())

	moveCount(up, 4)

	expected := `....H
....1
..432
.5...
T....
`
	assert.EqualString(expected, rope.String())
}
