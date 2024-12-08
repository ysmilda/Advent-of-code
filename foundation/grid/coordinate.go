package grid

import (
	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
)

// Coordinate represents a point in a 2D grid.
type Coordinate struct {
	X int
	Y int
}

// NewCoordinate creates a new Coordinate with the given x and y values.
func NewCoordinate(x, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

// Add returns the sum of two coordinates.
func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

// MoveInDirection returns the coordinate after moving in the given direction by the given increment.
func (c Coordinate) MoveInDirection(d Direction, increment uint) Coordinate {
	i := int(increment)
	switch d {
	case North:
		return c.Add(Coordinate{0, -i})

	case NorthEast:
		return c.Add(Coordinate{i, -i})

	case East:
		return c.Add(Coordinate{i, 0})

	case SouthEast:
		return c.Add(Coordinate{i, i})

	case South:
		return c.Add(Coordinate{0, i})

	case SouthWest:
		return c.Add(Coordinate{-i, i})

	case West:
		return c.Add(Coordinate{-i, 0})

	case NorthWest:
		return c.Add(Coordinate{-i, -i})
	}
	return c
}

// Subtract returns the difference of two coordinates.
func (c Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X - other.X,
		Y: c.Y - other.Y,
	}
}

// Multiply returns the product of two coordinates.
func (c Coordinate) Multiply(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X * other.X,
		Y: c.Y * other.Y,
	}
}

// Divide returns the division of two coordinates.
func (c Coordinate) Divide(other Coordinate) Coordinate {
	if c.X == 0 || other.X == 0 || c.Y == 0 || other.Y == 0 {
		panic("cannot divide by zero")
	}
	return Coordinate{
		X: c.X / other.X,
		Y: c.Y / other.Y,
	}
}

// Abs returns the absolute value of the coordinate.
func (c Coordinate) Abs() Coordinate {
	return Coordinate{
		X: aocmath.Abs(c.X),
		Y: aocmath.Abs(c.Y),
	}
}

// Invert returns a Coordinate which is the invert of the original.
func (c Coordinate) Invert() Coordinate {
	return Coordinate{
		X: -c.X,
		Y: -c.Y,
	}
}

// Within returns true if the given coordinate fits within the given grid size.
func (c Coordinate) Within(size Coordinate) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < size.X && c.Y < size.Y
}

// ManhattanDistance returns the Manhattan distance between two coordinates.
func (c Coordinate) ManhattanDistance(other Coordinate) int {
	return aocmath.Abs(c.X-other.X) + aocmath.Abs(c.Y-other.Y)
}
