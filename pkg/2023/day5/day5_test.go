package aoc2023day5

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testinput.txt
var testInputFile string

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 5, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInputFile),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 35, result)
}

func Test2023Day1Part2(t *testing.T) {
	solver := puzzle{
		input: parse(testInputFile),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 46, result)
}
