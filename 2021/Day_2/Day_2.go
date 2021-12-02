package day_2

import (
	solver "advent-of-code-2021/SolverInterface"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []step
}

type step struct {
	direction string
	amount    int
}

type position struct {
	x, y int
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	input, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{
		Input: input,
	}, nil
}

func (s Solver) GetDay() int {
	return 2
}

func (s Solver) Part1() (int, error) {
	current := position{0, 0}

	for _, step := range s.Input {
		switch step.direction {
		case "forward":
			current.y += step.amount

		case "down":
			current.x += step.amount

		case "up":
			current.x -= step.amount

		}
	}

	return current.x * current.y, nil
}

func (s Solver) Part2() (int, error) {
	current := position{0, 0}
	aim := 0

	for _, step := range s.Input {
		switch step.direction {
		case "forward":
			current.y += step.amount
			current.x += aim * step.amount

		case "down":
			aim += step.amount

		case "up":
			aim -= step.amount

		}
	}

	return current.x * current.y, nil
}

func parseInput() (output []step, err error) {
	lines := strings.Split(inputFile, "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")

		value, err := strconv.Atoi(split[1])

		if err != nil {
			return nil, err
		}

		output = append(output, step{
			direction: split[0],
			amount:    value,
		})
	}

	return output, err
}
