package day_2

import (
	_ "embed"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []Move
}

type Move struct {
	Input  int
	Output int
}

func (m Move) outcome1() int {
	// 1 = rock
	// 2 = paper
	// 3 = scissor

	if (m.Input == 1 && m.Output == 2) || (m.Input == 2 && m.Output == 3) || (m.Input == 3 && m.Output == 1) {
		return m.Output + 6
	}

	if m.Input == m.Output {
		return m.Output + 3
	}

	return m.Output
}

func (m Move) outcome2() int {
	// 1 = lose
	// 2 = draw
	// 3 = win

	switch m.Output {
	case 1:
		count := m.Input - 1
		if count < 1 {
			count = 3
		}
		return count
	case 2:
		return m.Input + 3
	case 3:
		count := m.Input + 1
		if count > 3 {
			count = 1
		}
		return 6 + count
	}

	return 0
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

func (s Solver) Part1() (string, error) {
	count := 0
	for _, move := range s.Input {
		count += move.outcome1()
	}
	return strconv.Itoa(count), nil
}

func (s Solver) Part2() (string, error) {
	count := 0
	for _, move := range s.Input {
		count += move.outcome2()
	}
	return strconv.Itoa(count), nil
}

func parseInput() ([]Move, error) {
	lines := strings.Split(inputFile, "\n")

	output := make([]Move, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, " ")

		switch splitLine[0] {
		case "A":
			output[i].Input = 1
		case "B":
			output[i].Input = 2
		case "C":
			output[i].Input = 3
		}

		switch splitLine[1] {
		case "X":
			output[i].Output = 1
		case "Y":
			output[i].Output = 2
		case "Z":
			output[i].Output = 3
		}
	}

	return output, nil
}
