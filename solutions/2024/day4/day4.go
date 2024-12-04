package aoc2024day4

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input grid.Grid[byte]
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 4
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	word := []byte("XMAS")

	for y := range s.input.GetHeight() {
		for x := range s.input.GetWidth() {
			start := grid.NewCoordinate(x, y)

			if s.input.Get(start) != word[0] {
				continue
			}

			for _, d := range []grid.Direction{
				grid.North,
				grid.NorthEast,
				grid.East,
				grid.SouthEast,
				grid.South,
				grid.SouthWest,
				grid.West,
				grid.NorthWest,
			} {
				found := true
				step := start
				for _, c := range word[1:] {
					step = step.MoveInDirection(d, 1)
					if !s.input.Valid(step) || s.input.Get(step) != c {
						found = false
						break
					}
				}

				if found {
					sum++
				}
			}
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for y := range s.input.GetHeight() {
		for x := range s.input.GetWidth() {
			start := grid.NewCoordinate(x, y)

			if s.input.Get(start) != 'A' {
				continue
			}

			step := start.MoveInDirection(grid.NorthEast, 1)
			if !s.input.Valid(step) {
				continue
			}

			switch s.input.Get(step) {
			default:
				continue

			case 'M':
				step = start.MoveInDirection(grid.SouthWest, 1)
				if !s.input.Valid(step) {
					continue
				}
				if s.input.Get(step) != 'S' {
					continue
				}

			case 'S':
				step = start.MoveInDirection(grid.SouthWest, 1)
				if !s.input.Valid(step) {
					continue
				}
				if s.input.Get(step) != 'M' {
					continue
				}
			}

			step = start.MoveInDirection(grid.NorthWest, 1)
			if !s.input.Valid(step) {
				continue
			}

			switch s.input.Get(step) {
			default:
				continue

			case 'M':
				step = start.MoveInDirection(grid.SouthEast, 1)
				if !s.input.Valid(step) {
					continue
				}
				if s.input.Get(step) != 'S' {
					continue
				}

			case 'S':
				step = start.MoveInDirection(grid.SouthEast, 1)
				if !s.input.Valid(step) {
					continue
				}
				if s.input.Get(step) != 'M' {
					continue
				}
			}

			sum++
		}
	}

	return sum, nil
}

func parse(input string) grid.Grid[byte] {
	lines := strings.Split(input, "\n")
	g := grid.NewGrid[byte](uint(len(lines[0])), uint(len(lines)))

	for i, line := range lines {
		g[i] = []byte(line)
	}

	return g
}
