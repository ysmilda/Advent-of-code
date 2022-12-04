package day_3

import (
	_ "embed"
	"errors"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	Input []Rucksack
}

type RuckSackGroup []Rucksack

func (r RuckSackGroup) FindCommon() rune {
	if len(r) < 3 {
		panic("Too little entries")
	}
	for _, char := range r[0].Combined {
		if strings.ContainsRune(r[1].Combined, char) && strings.ContainsRune(r[2].Combined, char) {
			return char
		}
	}
	return '\000'
}

type Rucksack struct {
	Combined     string
	Compartment1 string
	Compartment2 string
}

func (r Rucksack) FindCommon() rune {
	for _, char := range r.Compartment1 {
		if strings.ContainsRune(r.Compartment2, char) {
			return char
		}
	}
	return '\000'
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
	return 3
}

func (s Solver) Part1() (int, error) {
	count := 0
	for _, rucksack := range s.Input {
		common := rucksack.FindCommon()
		if common == '\000' {
			return 0, errors.New("No common letter found")
		}

		count += asciiToPriority(common)
	}
	return count, nil
}

func (s Solver) Part2() (int, error) {
	count := 0
	for i := 0; i < len(s.Input); i += 3 {
		common := RuckSackGroup(s.Input[i : i+3]).FindCommon()
		if common == '\000' {
			return 0, errors.New("No common letter found")
		}

		count += asciiToPriority(common)
	}
	return count, nil
}

func parseInput() ([]Rucksack, error) {
	lines := strings.Split(inputFile, "\n")

	output := make([]Rucksack, len(lines))
	for i, line := range lines {
		output[i] = Rucksack{
			Combined:     line,
			Compartment1: line[:len(line)/2],
			Compartment2: line[len(line)/2:],
		}
	}

	return output, nil
}

func asciiToPriority(input rune) int {
	if input >= 65 && input <= 90 {
		return int(input) - 65 + 27
	}
	if input >= 97 && input <= 122 {
		return int(input) - 97 + 1
	}
	return 0
}
