package aoc2023day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 3, day)
}

func TestPart1(t *testing.T) {
	partNumbers, symbols := parse(testInput)
	solver := puzzle{
		partNumbers: partNumbers,
		symbols:     symbols,
	}

	result, _ := solver.Part1()
	assert.Equal(t, 4361, result)
}

func TestPart2(t *testing.T) {
	partNumbers, symbols := parse(testInput)
	solver := puzzle{
		partNumbers: partNumbers,
		symbols:     symbols,
	}

	result, _ := solver.Part2()
	assert.Equal(t, 467835, result)
}
