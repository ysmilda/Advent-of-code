package grid

import (
	"errors"
)

// Grid represents a 2D grid of elements.
type Grid[T comparable] [][]T

// NewGrid creates a new grid with the given width and height.
func NewGrid[T comparable](width, height uint) Grid[T] {
	grid := make([][]T, height)

	for i := range grid {
		grid[i] = make([]T, width)
	}

	return grid
}

func CopyFrom[T comparable](in Grid[T]) Grid[T] {
	g := NewGrid[T](uint(in.GetWidth()), uint(in.GetHeight()))
	for i, row := range in {
		copy(g[i], row)
	}
	return g
}

// Valid returns true if the given coordinate is within the grid.
func (g Grid[T]) Valid(c Coordinate) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < g.GetWidth() && c.Y < g.GetHeight()
}

// GetWidth returns the width of the grid.
func (g Grid[T]) GetWidth() int {
	return len(g[0])
}

// GetHeight returns the height of the grid.
func (g Grid[T]) GetHeight() int {
	return len(g)
}

// Get returns the element at the given coordinate. It assumes the coordinate is valid.
func (g Grid[T]) Get(c Coordinate) T {
	return g[uint(c.Y)][uint(c.X)]
}

// Set sets the element at the given coordinate. It assumes the coordinate is valid.
func (g Grid[T]) Set(c Coordinate, b T) {
	g[uint(c.Y)][uint(c.X)] = b
}

// GetRow returns the row at the given y-coordinate.
func (g Grid[T]) GetRow(y uint) []T {
	return g[y]
}

// AddRow adds a row at the given index.
func (g Grid[T]) AddRow(row []T, index uint) Grid[T] {
	g = append(g, []T{})
	copy(g[index+1:], g[index:])
	g[index] = row
	return g
}

// GetColumn returns the column at the given x-coordinate.
func (g Grid[T]) GetColumn(x uint) []T {
	column := make([]T, len(g))

	for i, row := range g {
		column[i] = row[x]
	}

	return column
}

// AddColumn adds a column at the given index.
func (g Grid[T]) AddColumn(column []T, index uint) Grid[T] {
	for i, row := range g {
		entry := column[i]
		g[i] = append(row, entry)
		copy(g[i][index+1:], g[i][index:])
		g[i][index] = entry
	}
	return g
}

// Find returns the coordinate of the first occurrence of the given element.
func (g Grid[T]) Find(b T) (Coordinate, error) {
	for y, row := range g {
		for x, char := range row {
			if char == b {
				return NewCoordinate(x, y), nil
			}
		}
	}

	return Coordinate{}, errors.New("unable to find")
}

// FindAll returns the coordinates of all occurrences of the given element.
func (g Grid[T]) FindAll(b T) []Coordinate {
	var coordinates []Coordinate

	for y, row := range g {
		for x, char := range row {
			if char == b {
				coordinates = append(coordinates, NewCoordinate(x, y))
			}
		}
	}

	return coordinates
}
