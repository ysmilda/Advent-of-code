package aoc2023day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 11, day)
}

func TestCalculateDistances(t *testing.T) {
	testCases := []struct {
		step     int
		expected int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}

	solver := puzzle{
		grid: parse(testInput),
	}

	for _, tc := range testCases {
		result := solver.CalculateDistances(tc.step)
		assert.Equal(t, tc.expected, result)
	}
}
