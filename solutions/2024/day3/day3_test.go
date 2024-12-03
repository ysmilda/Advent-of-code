package aoc2024day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 3, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 161, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 48, result)
}
