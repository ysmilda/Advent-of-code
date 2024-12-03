package aoc2024day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 2, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		reports: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 2, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		reports: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 4, result)
}
