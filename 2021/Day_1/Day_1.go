package day_1

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputDay1 string

type Solver struct {
	Input []int
}

func GetSolver() (Solver, error) {
	input, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{
		Input: input,
	}, nil
}

func (d Solver) GetDay() int {
	return 1
}

func (d Solver) Part1() (count int) {
	for i := 1; i < len(d.Input); i++ {
		if d.Input[i] > d.Input[i-1] {
			count++
		}
	}

	return count
}

func (d Solver) Part2() (count int) {
	for i := 0; i < len(d.Input)-3; i++ {
		averageA := d.Input[i] + d.Input[i+1] + d.Input[i+2]
		averageB := d.Input[i+1] + d.Input[i+2] + d.Input[i+3]

		if averageB > averageA {
			count++
		}
	}

	return count
}

func parseInput() (output []int, err error) {
	lines := strings.Split(inputDay1, "\n")

	for _, line := range lines {
		value, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		output = append(output, value)
	}

	return output, err
}
