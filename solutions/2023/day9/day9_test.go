package aoc2023day9

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 9, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 114, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 2, result)
}
