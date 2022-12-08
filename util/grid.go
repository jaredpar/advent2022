package util

import (
	"errors"
	"fmt"
)

type Grid[T any] struct {
	values       []T
	columnLength int
}

func rowColumn(index, columnLength int) (row, column int) {
	column = index % columnLength
	row = (index - column) / columnLength
	return
}

func index(row, column, columnLength int) int {
	return (row * columnLength) + column
}

func NewGrid[T any](row, column int) *Grid[T] {
	if row == 0 || column == 0 {
		panic("row and column must be above 0")
	}

	values := make([]T, row*column)
	return &Grid[T]{values: values, columnLength: column}
}

func (g *Grid[T]) Count() int {
	return len(g.values)
}

func (g *Grid[T]) Index(row, column int) int {
	return index(row, column, g.columnLength)
}

func (g *Grid[T]) RowColumn(index int) (row, column int) {
	return rowColumn(index, g.columnLength)
}

func (g *Grid[T]) Rows() int {
	return len(g.values) / g.columnLength
}

func (g *Grid[T]) Columns() int {
	return g.columnLength
}

func (g *Grid[T]) Value(row, column int) T {
	index := g.Index(row, column)
	return g.values[index]
}

func (g *Grid[T]) SetValue(row, column int, value T) {
	index := g.Index(row, column)
	g.values[index] = value
}

func (g *Grid[T]) SetAll(value T) {
	SetAll(g.values, value)
}

func (g *Grid[T]) Resize(row, column int) {
	if row == g.Rows() && column == g.Columns() {
		return
	}

	oldColumnLength := g.columnLength
	oldValues := g.values

	g.values = make([]T, row*column)
	g.columnLength = column

	for i, v := range oldValues {
		r, c := rowColumn(i, oldColumnLength)
		if r < row && c < column {
			g.SetValue(r, c, v)
		}
	}
}

// Parse out a grid from a series of single digit entries on a
// line like the following
//
// 0123
// 4567
func ParseGridFromLines(lines []string) (*Grid[int], error) {
	if len(lines) == 0 {
		return nil, errors.New("need at least one line")
	}

	grid := NewGrid[int](len(lines), len(lines[0]))
	for row, line := range lines {
		if len(line) != grid.Rows() {
			return nil, fmt.Errorf("line has wrong length: %s", line)
		}

		for col, r := range line {
			digit, err := RuneToInt(r)
			if err != nil {
				return nil, err
			}

			grid.SetValue(row, col, digit)
		}
	}

	return grid, nil
}
