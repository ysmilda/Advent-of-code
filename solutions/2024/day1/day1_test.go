package aoc2024day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 1, day)
}

func TestPart1(t *testing.T) {
	left, right := parse(testInput)
	solver := puzzle{
		left:  left,
		right: right,
	}

	result, _ := solver.Part1()
	assert.Equal(t, 11, result)
}

func TestPart2(t *testing.T) {
	left, right := parse(testInput)
	solver := puzzle{
		left:  left,
		right: right,
	}

	result, _ := solver.Part2()
	assert.Equal(t, 31, result)
}
