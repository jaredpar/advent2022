package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
	"github.com/jaredpar/advent2022/util"
)

func TestPart1Sample(t *testing.T) {
	assert := testUtil.NewAssert(t)
	assert.EqualInt(5, getStartOffset("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.EqualInt(6, getStartOffset("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.EqualInt(11, getStartOffset("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	lines, err := util.ReadLines(f, "input.txt")
	assert.NotError(err)
	assert.EqualInt(1, len(lines))
	assert.EqualInt(1779, getStartOffset(lines[0]))
}
