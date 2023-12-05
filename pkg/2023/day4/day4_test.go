package aoc2023day4

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
	assert.Equal(t, 4, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		cards: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 13, result)
}

func Test2023Day1Part2(t *testing.T) {
	solver := puzzle{
		cards: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 30, result)
}
