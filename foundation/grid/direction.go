package grid

// Direction represents a direction in a 2D grid.
type Direction int

const (
	None Direction = iota
	North
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

// GetDirection returns the direction from start to end.
func GetDirection(start Coordinate, end Coordinate) Direction {
	if start.X == end.X {
		if start.Y > end.Y {
			return North
		} else if start.Y < end.Y {
			return South
		}
	} else if start.Y == end.Y {
		if start.X > end.X {
			return West
		} else if start.X < end.X {
			return East
		}
	} else if start.X < end.X {
		if start.Y < end.Y {
			return SouthWest
		} else if start.Y > end.Y {
			return NorthWest
		}
	} else if start.X > end.X {
		if start.Y < end.Y {
			return SouthEast
		} else if start.Y > end.Y {
			return NorthEast
		}
	}

	return None
}

func (d Direction) North() bool {
	return d == North
}

func (d Direction) NorthEast() bool {
	return d == NorthEast
}

func (d Direction) East() bool {
	return d == East
}

func (d Direction) SouthEast() bool {
	return d == SouthEast
}

func (d Direction) South() bool {
	return d == South
}

func (d Direction) SouthWest() bool {
	return d == SouthWest
}

func (d Direction) West() bool {
	return d == West
}

func (d Direction) NorthWest() bool {
	return d == NorthWest
}
