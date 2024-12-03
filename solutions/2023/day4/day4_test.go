package aoc2023day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

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

func TestPart2(t *testing.T) {
	solver := puzzle{
		cards: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 30, result)
}
