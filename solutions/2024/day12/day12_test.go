package aoc2024day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 12, day)
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 140},
		{
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			expected: 772,
		},
		{
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			expected: 1930,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			input: parse(tc.input),
		}

		result, _ := solver.Part1()
		assert.Equal(t, tc.expected, result)
	}

}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 80},
		{
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			expected: 436,
		},
		{
			input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			expected: 368,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			input: parse(tc.input),
		}

		result, _ := solver.Part2()
		assert.Equal(t, tc.expected, result)
	}
}
