package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
	"github.com/jaredpar/advent2022/util"
)

func TestPart1Sample(t *testing.T) {
	assert := testUtil.NewAssert(t)
	assert.EqualInt(5, getPacketOffset("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.EqualInt(6, getPacketOffset("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.EqualInt(11, getPacketOffset("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	lines, err := util.ReadLines(f, "input.txt")
	assert.NotError(err)
	assert.EqualInt(1, len(lines))
	assert.EqualInt(1779, getPacketOffset(lines[0]))
}

func TestPart2Sample(t *testing.T) {
	assert := testUtil.NewAssert(t)
	assert.EqualInt(23, getMessageOffset("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.EqualInt(23, getMessageOffset("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.EqualInt(26, getMessageOffset("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
