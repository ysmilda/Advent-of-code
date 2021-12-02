package day_1

import (
	solver "advent-of-code-2021/SolverInterface"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []int
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
	return 1
}

func (s Solver) Part1() (int, error) {
	count := 0

	for i := 1; i < len(s.Input); i++ {
		if s.Input[i] > s.Input[i-1] {
			count++
		}
	}

	return count, nil
}

func (s Solver) Part2() (int, error) {
	count := 0

	for i := 0; i < len(s.Input)-3; i++ {
		averageA := s.Input[i] + s.Input[i+1] + s.Input[i+2]
		averageB := s.Input[i+1] + s.Input[i+2] + s.Input[i+3]

		if averageB > averageA {
			count++
		}
	}

	return count, nil
}

func parseInput() (output []int, err error) {
	lines := strings.Split(inputFile, "\n")

	for _, line := range lines {
		value, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		output = append(output, value)
	}

	return output, err
}
