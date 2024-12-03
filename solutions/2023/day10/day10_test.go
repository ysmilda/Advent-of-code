package aoc2023day10

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 10, day)
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    ".....\n.S-7.\n.|.|.\n.L-J.\n.....",
			expected: 4,
		},
		{
			input:    "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...",
			expected: 8,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			grid: parse(tc.input),
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
		{
			input:    "...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n...........",
			expected: 4,
		},
		{
			input:    ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...",
			expected: 8,
		},
		{
			input:    "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L",
			expected: 10,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			grid: parse(tc.input),
		}

		result, _ := solver.Part2()
		assert.Equal(t, tc.expected, result)
	}
}
