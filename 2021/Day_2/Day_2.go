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
	Input []Step
}

type Step struct {
	Direction string
	Amount    int
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
		switch step.Direction {
		case "forward":
			current.y += step.Amount

		case "down":
			current.x += step.Amount

		case "up":
			current.x -= step.Amount

		}
	}

	return current.x * current.y, nil
}

func (s Solver) Part2() (int, error) {
	current := position{0, 0}
	aim := 0

	for _, step := range s.Input {
		switch step.Direction {
		case "forward":
			current.y += step.Amount
			current.x += aim * step.Amount

		case "down":
			aim += step.Amount

		case "up":
			aim -= step.Amount

		}
	}

	return current.x * current.y, nil
}

func parseInput() (output []Step, err error) {
	lines := strings.Split(inputFile, "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")

		value, err := strconv.Atoi(split[1])

		if err != nil {
			return nil, err
		}

		output = append(output, Step{
			Direction: split[0],
			Amount:    value,
		})
	}

	return output, err
}
