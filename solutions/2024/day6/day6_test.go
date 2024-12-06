package aoc2024day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 6, day)
}

func TestPart1(t *testing.T) {
	g, p, d := parse(testInput)
	solver := puzzle{
		grid:      g,
		start:     p,
		direction: d,
	}

	result, _ := solver.Part1()
	assert.Equal(t, 41, result)
}

func TestPart2(t *testing.T) {
	g, p, d := parse(testInput)
	solver := puzzle{
		grid:      g,
		start:     p,
		direction: d,
	}

	result, _ := solver.Part2()
	assert.Equal(t, 6, result)
}
