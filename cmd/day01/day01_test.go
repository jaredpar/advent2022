package main

import (
	"testing"
	"github.com/jaredpar/advent2022/testUtil"
)

func TestPart1(t *testing.T) {
	assert := testUtil.NewAssert(t)
	value := getSumOfMax(1)
	assert.EqualInt(69281, value)
}

func TestPart2(t *testing.T) {
	assert := testUtil.NewAssert(t)
	value := getSumOfMax(3)
	assert.EqualInt(201524, value)
}