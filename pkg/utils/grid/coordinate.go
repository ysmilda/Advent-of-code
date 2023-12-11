package grid

import (
	"github.com/ysmilda/Advent-of-code/pkg/utils/aocmath"
)

type Coordinate struct {
	X uint
	Y uint
}

func NewCoordinate(x, y uint) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

func (c Coordinate) AddDirection(d Direction, increment uint) Coordinate {
	switch d {
	case North:
		return c.Add(Coordinate{0, -increment})

	case South:
		return c.Add(Coordinate{0, increment})

	case West:
		return c.Add(Coordinate{-increment, 0})

	case East:
		return c.Add(Coordinate{increment, 0})
	}
	return c
}

func (c Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X - other.X,
		Y: c.Y - other.Y,
	}
}

func (c Coordinate) Multiply(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X * other.X,
		Y: c.Y * other.Y,
	}
}

func (c Coordinate) Divide(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X / other.X,
		Y: c.Y / other.Y,
	}
}

func (c Coordinate) Abs() Coordinate {
	return Coordinate{
		X: aocmath.Abs(c.X),
		Y: aocmath.Abs(c.Y),
	}
}

func (c Coordinate) ManhattanDistance(other Coordinate) uint {
	return aocmath.Abs(c.X-other.X) + aocmath.Abs(c.Y-other.Y)
}
