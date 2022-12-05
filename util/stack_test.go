package util

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestStackBasic(t *testing.T) {
	assert := testUtil.NewAssert(t)
	stack := NewStack[int]()
	stack.Push(13)

	assert.EqualInt(1, stack.Length())
	assert.EqualInt(13, stack.Pop())
	assert.EqualInt(0, stack.Length())
}

func TestStackCapacity(t *testing.T) {
	assert := testUtil.NewAssert(t)
	stack := NewStack[int]()

	for i := 1; i <= 10; i++ {
		stack.Push(i)
		assert.EqualInt(i, stack.Length())
	}

	next := 10
	for stack.Length() > 0 {
		assert.EqualInt(next, stack.Pop())
		next--
	}
}
