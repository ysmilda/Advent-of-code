package aoc2024day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 7, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 3749, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 11387, result)
}
