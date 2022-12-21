package util

type VirtualGrid[T any] struct {
	grid              *Grid[T]
	minRow, minColumn int
	defaultValue      T
	hasDefaultValue   bool
}

func NewVirtualGrid[T any]() *VirtualGrid[T] {
	grid := NewGrid[T](1, 1)
	var d T
	return &VirtualGrid[T]{grid, 0, 0, d, false}
}

func NewVirtualGridWithDefaultValue[T any](value T) *VirtualGrid[T] {
	grid := NewGrid[T](1, 1)
	grid.SetValue(0, 0, value)
	return &VirtualGrid[T]{grid, 0, 0, value, true}
}

func (grid *VirtualGrid[T]) Rows() int {
	return grid.grid.Rows()
}

func (grid *VirtualGrid[T]) MinRow() int {
	return grid.minRow
}

func (grid *VirtualGrid[T]) MaxRow() int {
	return grid.minRow + grid.grid.Rows()
}

func (grid *VirtualGrid[T]) Columns() int {
	return grid.grid.Columns()
}

func (grid *VirtualGrid[T]) MinColumn() int {
	return grid.minColumn
}

func (grid *VirtualGrid[T]) MaxColumn() int {
	return grid.minColumn + grid.grid.Rows()
}

func (grid *VirtualGrid[T]) SetValue(row, column int, value T) {
	if row < grid.minRow || row > grid.MaxRow() || column < grid.minColumn || column > grid.MaxColumn() {
		minRow := Min(row, grid.minRow)
		rows := Max(row+1, grid.MaxRow()) - minRow
		minColumn := Min(column, grid.minColumn)
		columns := Max(column+1, grid.MaxColumn()) - minColumn
		grid.resize(minRow, rows, minColumn, columns)
	}

	row = row + Abs(grid.minRow)
	column = column + Abs(grid.minColumn)
	grid.grid.SetValue(row, column, value)
}

func (grid *VirtualGrid[T]) Value(row, column int) T {
	row = row + Abs(grid.minRow)
	column = column + Abs(grid.minColumn)
	return grid.grid.Value(row, column)
}

func (grid *VirtualGrid[T]) resize(minRow, rows, minCol, cols int) {
	oldGrid := grid.grid
	newGrid := NewGrid[T](rows, cols)
	rowAdjust := grid.minRow - minRow
	colAdjust := grid.minColumn - minCol

	if grid.hasDefaultValue {
		newGrid.SetAll(grid.defaultValue)
	}

	for r := 0; r < oldGrid.Rows(); r++ {
		for c := 0; c < oldGrid.Columns(); c++ {
			value := oldGrid.Value(r, c)
			newGrid.SetValue(r+rowAdjust, c+colAdjust, value)
		}
	}

	grid.grid = newGrid
	grid.minRow = minRow
	grid.minColumn = minCol
}
