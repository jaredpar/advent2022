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
