package aoc2024day6

import (
	_ "embed"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	grid      grid.Grid[bool]
	start     grid.Coordinate
	direction grid.Direction
}

func MustGetSolver() solver.Solver {
	g, s, d := parse(inputFile)
	return puzzle{
		grid:      g,
		start:     s,
		direction: d,
	}
}

func (s puzzle) GetDay() int {
	return 6
}

func (s puzzle) Part1() (int, error) {
	visited := make(map[grid.Coordinate]struct{})
	visited[s.start] = struct{}{}

	Solve(s.grid, s.start, s.direction,
		func(p grid.Coordinate, _ grid.Direction) bool {
			return s.grid.Valid(p)
		},
		func(p grid.Coordinate) {
			visited[p] = struct{}{}
		},
	)

	return len(visited), nil
}

type move struct {
	position  grid.Coordinate
	direction grid.Direction
}

func (s puzzle) Part2() (int, error) {
	sum := atomic.Int64{}
	wg := sync.WaitGroup{}

	for y := range s.grid.GetHeight() {
		for x := range s.grid.GetWidth() {
			coord := grid.NewCoordinate(x, y)
			if coord == s.start || s.grid.Get(coord) {
				continue
			}

			g := grid.CopyFrom(s.grid)
			g.Set(coord, true)
			wg.Add(1)

			// Speed up bruteforce by using all the cores.
			go func(g grid.Grid[bool], start grid.Coordinate, direction grid.Direction) {
				visited := make(map[move]struct{})
				Solve(g, start, direction,
					func(p grid.Coordinate, d grid.Direction) bool {
						if !g.Valid(p) {
							return false
						}

						m := move{p, d}
						// Detect if we are in a loop, if so add one to the sum and return false. Else continue.
						// If the coordinate + direction matches one we've seen before we've entered a loop.
						if _, exists := visited[m]; exists {
							sum.Add(1)
							return false
						}

						visited[m] = struct{}{}
						return true

					}, func(p grid.Coordinate) {},
				)

				wg.Done()
			}(g, s.start, s.direction)
		}
	}

	wg.Wait()

	return int(sum.Load()), nil
}

func Solve(g grid.Grid[bool], start grid.Coordinate, direction grid.Direction, valid func(grid.Coordinate, grid.Direction) bool, f func(grid.Coordinate)) {
	position := start
	for {
		next := position.MoveInDirection(direction, 1)
		if !valid(next, direction) {
			break
		}

		if g.Get(next) { // There is a blockade, we need to rotate right.
			direction = direction.RotateRight()
			next = position.MoveInDirection(direction, 1) // Move to the next position.

			if g.Get(next) { // This might also be an obstacle, if so rotate again, we can always move backwards.
				direction = direction.RotateRight()
				next = position.MoveInDirection(direction, 1)
			}
		}

		position = next
		f(position)
	}
}

func parse(input string) (grid.Grid[bool], grid.Coordinate, grid.Direction) {
	var (
		lines     = strings.Split(input, "\n")
		g         = grid.NewGrid[bool](uint(len(lines[0])), uint(len(lines)))
		start     grid.Coordinate
		direction grid.Direction
	)

	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				start = grid.NewCoordinate(x, y)
				direction = grid.Up
				continue
			}
			g.Set(grid.NewCoordinate(x, y), char == '#')
		}
	}

	return g, start, direction
}
