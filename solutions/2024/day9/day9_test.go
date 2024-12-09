package aoc2024day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `2333133121414131402`

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
	assert.Equal(t, 1928, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 2858, result)
}
