package aoc2023day11

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testinput.txt
var testInput string

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 11, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		grid: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 374, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		grid: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 0, result)
}
