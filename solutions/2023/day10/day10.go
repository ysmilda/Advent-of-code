package aoc2023day10

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	grid grid.Grid[byte]
}

func (s puzzle) Next(previous grid.Coordinate, current grid.Coordinate) grid.Coordinate {
	d := grid.GetDirection(previous, current)
	pipe := s.grid.GetCoordinate(current)
	switch pipe {
	case '|':
		if d.North() || d.South() {
			return current.AddDirection(d, 1)
		}

	case '-':
		if d.West() || d.East() {
			return current.AddDirection(d, 1)
		}

	case 'L':
		switch d {
		case grid.West:
			return current.AddDirection(grid.North, 1)
		case grid.South:
			return current.AddDirection(grid.East, 1)
		}

	case 'F':
		switch d {
		case grid.West:
			return current.AddDirection(grid.South, 1)
		case grid.North:
			return current.AddDirection(grid.East, 1)
		}

	case 'J':
		switch d {
		case grid.East:
			return current.AddDirection(grid.North, 1)
		case grid.South:
			return current.AddDirection(grid.West, 1)
		}

	case '7':
		switch d {
		case grid.East:
			return current.AddDirection(grid.South, 1)
		case grid.North:
			return current.AddDirection(grid.West, 1)
		}

	case 'S':
		next := grid.Coordinate{}
		validate := func(x, y int) bool {
			next = grid.NewCoordinate(x, y)
			if s.grid.Valid(next) {
				d := grid.GetDirection(current, next)
				switch s.grid.GetCoordinate(next) {
				case '|':
					return d.North() || d.South()

				case '-':
					return d.West() || d.East()

				case 'L':
					return d.West() || d.South()

				case 'F':
					return d.West() || d.North()

				case 'J':
					return d.East() || d.South()

				case '7':
					return d.East() || d.North()
				}
			}
			return false
		}

		if validate(current.X-1, current.Y) {
			return next
		}
		if validate(current.X+1, current.Y) {
			return next
		}
		if validate(current.X, current.Y-1) {
			return next
		}
		if validate(current.X, current.Y+1) {
			return next
		}
	}

	panic("no pipe found or invalid direction ")
}

func (s puzzle) Loop() int {
	start, err := s.grid.Find('S')
	if err != nil {
		panic("unable to find start")
	}
	previous, next := start, start
	steps := 1
	for {
		next, previous = s.Next(previous, next), next
		if next == start {
			break
		}
		steps++
	}
	return steps
}

func MustGetSolver() solver.Solver {
	return puzzle{
		grid: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 10
}

func (s puzzle) Part1() (int, error) {
	steps := s.Loop()
	return steps / 2, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0
	_ = s.Loop()
	return sum, nil
}

func parse(input string) grid.Grid[byte] {
	lines := strings.Split(input, "\n")
	output := grid.NewGrid[byte](uint(len(lines[0])), uint(len(lines)))

	for i, line := range lines {
		output[i] = []byte(line)
	}

	return output
}
