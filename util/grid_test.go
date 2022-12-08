package util

import (
	"testing"

	"github.com/jaredpar/advent2022/testUtil"
)

func TestRowColumn(t *testing.T) {
	assert := testUtil.NewAssert(t)
	g := NewGrid[int](10, 10)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			index := g.Index(r, c)
			actualRow, actualCol := g.RowColumn(index)
			assert.EqualInt(r, actualRow)
			assert.EqualInt(c, actualCol)
		}
	}
}

func TestExpand(t *testing.T) {
	g := NewGrid[int](5, 5)
	for i := 0; i < g.Count(); i++ {
		g.values[i] = i
	}

	assert := testUtil.NewAssert(t)
	g.Resize(10, 10)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			if r < 5 && c < 5 {
				value := index(r, c, 5)
				assert.EqualInt(value, g.Value(r, c))
			} else {
				assert.EqualInt(0, g.Value(r, c))
			}
		}
	}
}

func TestShrink(t *testing.T) {
	g := NewGrid[int](5, 5)
	for i := 0; i < g.Count(); i++ {
		g.values[i] = 1
	}

	assert := testUtil.NewAssert(t)
	g.Resize(3, 3)
	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Columns(); c++ {
			assert.EqualInt(1, g.Value(r, c))
		}
	}
}
