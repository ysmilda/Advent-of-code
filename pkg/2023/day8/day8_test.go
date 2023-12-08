package aoc2023day8

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 8, day)
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input  string
		result int
	}{
		{
			`RL

			AAA = (BBB, CCC)
			BBB = (DDD, EEE)
			CCC = (ZZZ, GGG)
			DDD = (DDD, DDD)
			EEE = (EEE, EEE)
			GGG = (GGG, GGG)
			ZZZ = (ZZZ, ZZZ)`,
			2,
		},
		{
			`LLR

			AAA = (BBB, BBB)
			BBB = (AAA, ZZZ)
			ZZZ = (ZZZ, ZZZ)`,
			6,
		},
	}

	for _, tc := range testCases {
		instructions, mapping := parse(tc.input)
		solver := puzzle{
			instructions: instructions,
			directions:   mapping,
		}

		result, _ := solver.Part1()
		assert.Equal(t, tc.result, result)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input  string
		result int
	}{
		{
			`LR

			11A = (11B, XXX)
			11B = (XXX, 11Z)
			11Z = (11B, XXX)
			22A = (22B, XXX)
			22B = (22C, 22C)
			22C = (22Z, 22Z)
			22Z = (22B, 22B)
			XXX = (XXX, XXX)`,
			6,
		},
	}

	for _, tc := range testCases {
		instructions, mapping := parse(tc.input)
		solver := puzzle{
			instructions: instructions,
			directions:   mapping,
		}

		result, _ := solver.Part2()
		assert.Equal(t, tc.result, result)
	}
}
