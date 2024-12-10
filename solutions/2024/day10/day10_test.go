package aoc2024day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 10, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 36, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 81, result)
}
