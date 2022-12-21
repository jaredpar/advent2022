package util

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestVirtualGridSimple(t *testing.T) {
	assert := testUtil.NewAssert(t)

	grid := NewVirtualGrid[int]()
	assert.EqualInt(0, grid.Value(0, 0))

	grid.SetValue(2, 2, 42)
	assert.EqualInt(42, grid.Value(2, 2))
	assert.EqualInt(3, grid.grid.Rows())
	assert.EqualInt(3, grid.grid.Columns())

	grid.SetValue(-1, -2, 13)
	assert.EqualInt(13, grid.Value(-1, -2))
	assert.EqualInt(4, grid.grid.Rows())
	assert.EqualInt(5, grid.grid.Columns())
}

func TestVirtualGridWithDefaultValue(t *testing.T) {
	assert := testUtil.NewAssert(t)

	grid := NewVirtualGridWithDefaultValue('.')
	assert.EqualRune('.', grid.Value(0, 0))

	grid.SetValue(13, 13, 'r')
	assert.EqualRune('r', grid.Value(13, 13))
	for i := 0; i < 10; i++ {
		assert.EqualRune('.', grid.Value(i, i))
	}
}
