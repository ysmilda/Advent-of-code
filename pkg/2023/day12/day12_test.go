package aoc2023day12

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 12, day)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    string
		expected int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
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
	t.Parallel()

	testCases := []struct {
		input    string
		expected int
	}{}

	for _, tc := range testCases {
		solver := puzzle{
			input: parse(tc.input),
		}

		result, _ := solver.Part2()
		assert.Equal(t, tc.expected, result)
	}
}
