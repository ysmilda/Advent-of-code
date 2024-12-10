package aoc2024day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 8, day)
}

func TestPart1(t *testing.T) {
	antennas, size := parse(testInput)
	solver := puzzle{
		antennas: antennas,
		size:     size,
	}

	result, _ := solver.Part1()
	assert.Equal(t, 14, result)
}

func TestPart2(t *testing.T) {
	antennas, size := parse(testInput)
	solver := puzzle{
		antennas: antennas,
		size:     size,
	}

	result, _ := solver.Part2()
	assert.Equal(t, 34, result)
}
