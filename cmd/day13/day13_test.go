package main

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestCompareSingle(t *testing.T) {
	assert := testUtil.NewAssert(t)

	one := newPacketSingle(1)
	three := newPacketSingle(3)

	assert.True(one.compare(three) < 0)
	assert.True(three.compare(one) > 0)
	assert.True(three.compare(three) == 0)
}

func TestCompareList(t *testing.T) {
	assert := testUtil.NewAssert(t)

	one := newPacketSingle(1)
	oneList := newPacketList([]packet{one})
	three := newPacketSingle(3)

	assert.True(one.compare(oneList) == 0)
	assert.True(oneList.compare(oneList) == 0)
	assert.True(oneList.compare(one) == 0)
	assert.True(oneList.compare(three) < 0)
}

func TestRoundTrip(t *testing.T) {
	assert := testUtil.NewAssert(t)

	core := func(line string) {
		p := parsePacket(line)
		assert.EqualString(line, p.String())
	}

	core("[]")
	core("[1]")
	core("[1,2]")
	core("[1,[2]]")
	core("[1,[2,3]]")
	core("[[1],[2,3,4]]")
}

func TestCompareFun(t *testing.T) {
	assert := testUtil.NewAssert(t)
	p1 := parsePacket("[[1],[2,3,4]]")
	p2 := parsePacket("[[1],4]")
	assert.True(p1.compare(p2) < 0)
}
