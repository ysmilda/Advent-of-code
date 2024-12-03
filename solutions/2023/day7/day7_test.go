package aoc2023day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 7, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: testInput,
	}

	result, _ := solver.Part1()
	assert.Equal(t, 6440, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: testInput,
	}

	result, _ := solver.Part2()
	assert.Equal(t, 5905, result)
}
