package aoc2024day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47

`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 5, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 143, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 123, result)
}
