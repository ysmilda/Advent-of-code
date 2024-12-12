package aoc2024day12

import (
	_ "embed"
	"maps"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input grid.Grid[int]
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 12
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	input := grid.CopyFrom(s.input)
	visited := map[grid.Coordinate]bool{}
	index := 255

	for y, row := range input {
		for x := range row {
			coord := grid.NewCoordinate(x, y)
			if _, ok := visited[coord]; ok {
				continue
			}

			area, perimeter := findAreaAndPerimeter(coord, index, input)

			// Check for every perimeter coordinate which edges are adjacent to the original area.
			per := 0
			for coord := range perimeter {
				for _, direction := range grid.CardinalDirections {
					if edge := coord.MoveInDirection(direction, 1); input.Valid(edge) && input.Get(edge) == index {
						per++
					}
				}
			}

			sum += len(area) * per
			index++
			maps.Insert(visited, maps.All(area))
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	return sum, nil
}

func parse(input string) grid.Grid[int] {
	return grid.NewGridFrom(strings.Split(input, "\n"), func(r rune) int {
		return int(byte(r))
	})
}

func check(start grid.Coordinate, input grid.Grid[int], c int, valid func(grid.Coordinate), invalid func(grid.Coordinate), visited func(grid.Coordinate) bool) {
	if visited(start) {
		return
	}

	if !input.Valid(start) || input.Get(start) != c {
		invalid(start)
		return
	}

	valid(start)
	for _, direction := range grid.CardinalDirections {
		check(start.MoveInDirection(direction, 1), input, c, valid, invalid, visited)
	}
}

func findAreaAndPerimeter(start grid.Coordinate, index int, input grid.Grid[int]) (area map[grid.Coordinate]bool, perimeter map[grid.Coordinate]bool) {
	area, perimeter = map[grid.Coordinate]bool{}, map[grid.Coordinate]bool{}

	// flood fill
	check(start, input, input.Get(start),
		func(c grid.Coordinate) {
			area[c] = true
			input.Set(c, index) // Give the coordinate a new value to distinguish between various regions with the same value.
		},
		func(c grid.Coordinate) {
			perimeter[c] = true
		},
		func(c grid.Coordinate) bool {
			_, ok := area[c]
			return ok
		},
	)

	return
}
