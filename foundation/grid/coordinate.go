package grid

import (
	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
)

type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) Coordinate {
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
	if c.X == 0 || other.X == 0 || c.Y == 0 || other.Y == 0 {
		panic("cannot divide by zero")
	}
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

func (c Coordinate) ManhattanDistance(other Coordinate) int {
	return aocmath.Abs(c.X-other.X) + aocmath.Abs(c.Y-other.Y)
}
