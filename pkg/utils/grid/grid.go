package grid

import (
	"errors"
)

type Grid[T comparable] [][]T

func NewGrid[T comparable](width, height uint) Grid[T] {
	grid := make([][]T, height)

	for i := range grid {
		grid[i] = make([]T, width)
	}

	return grid
}

func (g Grid[T]) Valid(c Coordinate) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < g.GetWidth() && c.Y < g.GetHeight()
}

func (g Grid[T]) GetWidth() uint {
	return uint(len(g[0]))
}

func (g Grid[T]) GetHeight() uint {
	return uint(len(g))
}

func (g Grid[T]) Get(x, y uint) T {
	return g[y][x]
}

func (g Grid[T]) Set(x, y uint, b T) {
	g[y][x] = b
}

func (g Grid[T]) GetCoordinate(c Coordinate) T {
	return g.Get(c.X, c.Y)
}

func (g Grid[T]) SetCoordinate(c Coordinate, b T) {
	g.Set(c.X, c.Y, b)
}

func (g Grid[T]) GetRow(y uint) []T {
	return g[y]
}

func (g Grid[T]) GetColumn(x uint) []T {
	column := make([]T, len(g))

	for i, row := range g {
		column[i] = row[x]
	}

	return column
}

func (g Grid[T]) Find(b T) (Coordinate, error) {
	for y, row := range g {
		for x, char := range row {
			if char == b {
				return NewCoordinate(uint(x), uint(y)), nil
			}
		}
	}

	return Coordinate{}, errors.New("unable to find")
}

func (g Grid[T]) FindAll(b T) []Coordinate {
	var coordinates []Coordinate

	for y, row := range g {
		for x, char := range row {
			if char == b {
				coordinates = append(coordinates, NewCoordinate(uint(x), uint(y)))
			}
		}
	}

	return coordinates
}
