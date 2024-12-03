package aoc2023day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 1, day)
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input   string
		outcome int
	}{
		{
			input:   "1abc2",
			outcome: 12,
		},
		{
			input:   "pqr3stu8vwx",
			outcome: 38,
		},
		{
			input:   "a1b2c3d4e5f",
			outcome: 15,
		},
		{
			input:   "treb7uchet",
			outcome: 77,
		},
		{
			input:   "asc9bdgbds9",
			outcome: 99,
		},
		{
			input:   "33291six",
			outcome: 31,
		},
		{
			input:   "0asds1asf2",
			outcome: 2,
		},
		{
			input:   "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			outcome: 142,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			input: parse(tc.input),
		}

		result, _ := solver.Part1()
		assert.Equal(t, tc.outcome, result)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input   string
		outcome int
	}{
		{
			input:   "two1nine",
			outcome: 29,
		},
		{
			input:   "eightwothree",
			outcome: 83,
		},
		{
			input:   "abcone2threexyz",
			outcome: 13,
		},
		{
			input:   "xtwone3four",
			outcome: 24,
		},
		{
			input:   "4nineeightseven2",
			outcome: 42,
		},
		{
			input:   "zoneight234",
			outcome: 14,
		},
		{
			input:   "7pqrstsixteen",
			outcome: 76,
		},
		{
			input:   "one2three\nfour5six\nseven8nine",
			outcome: 138,
		},
	}

	for _, tc := range testCases {
		solver := puzzle{
			input: parse(tc.input),
		}

		result, _ := solver.Part2()
		assert.Equal(t, tc.outcome, result)
	}
}
