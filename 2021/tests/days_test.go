package main

import (
	day_1 "advent-of-code-2021/Day_1"
	day_2 "advent-of-code-2021/Day_2"
	solver "advent-of-code-2021/SolverInterface"
	"testing"
)

type day struct {
	tests []input
}

type input struct {
	solve    solver.SolverInterface
	outcome1 int
	outcome2 int
}

var days []day = []day{
	{tests: []input{{
		solve: day_1.Solver{
			Input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
		},
		outcome1: 7,
		outcome2: 5,
	}}},
	{tests: []input{{
		solve: day_2.Solver{
			Input: []day_2.Step{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}},
		},
		outcome1: 150,
		outcome2: 900,
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
