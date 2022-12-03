package main

import (
	"testing"

	day_1 "advent-of-code-2022/Day_1"
	solver "advent-of-code-2022/SolverInterface"
)

type day struct {
	tests []input
}

type input struct {
	solve    solver.SolverInterface
	outcome1 int
	outcome2 int
}

// TODO: Adapt tests to accept the string input.
var days []day = []day{
	{tests: []input{{
		solve: day_1.Solver{
			Input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
		},
		outcome1: 7,
		outcome2: 5,
	}}},
}

func TestPart1(t *testing.T) {
	for dayIndex, input := range days {
		for testIndex, test := range input.tests {
			outcome, err := test.solve.Part1()
			if err != nil {
				t.Errorf("Failed to solve test %d for part 1 of day %d: %s", testIndex+1, dayIndex+1, err.Error())
			}

			if outcome != test.outcome1 {
				t.Errorf("Failed to solve test %d for part 1 of day %d. Expected %d, but got %d", testIndex+1, dayIndex+1, test.outcome1, outcome)
			}
		}
	}
}

func TestPart2(t *testing.T) {
	for dayIndex, input := range days {
		for testIndex, test := range input.tests {
			outcome, err := test.solve.Part2()
			if err != nil {
				t.Errorf("Failed to solve test %d for part 2 of day %d: %s", testIndex+1, dayIndex+1, err.Error())
			}

			if outcome != test.outcome2 {
				t.Errorf("Failed to solve test %d for part 2 of day %d. Expected %d, but got %d", testIndex+1, dayIndex+1, test.outcome2, outcome)
			}
		}
	}
}
