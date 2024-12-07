package aoc2024day7

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocslices"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type equation struct {
	result  int
	numbers []int
}

type puzzle struct {
	input []equation
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 7
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, equation := range s.input {
		for _, operators := range aocslices.Iterate(2, len(equation.numbers)-1) {
			res := equation.numbers[0]
			for j, number := range equation.numbers[1:] {
				switch operators[j] {
				case 0:
					res += number
				case 1:
					res *= number
				}
			}
			if res == equation.result {
				sum += equation.result
				break
			}
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, equation := range s.input {
		for _, operators := range aocslices.Iterate(3, len(equation.numbers)-1) {
			res := equation.numbers[0]
			for j, number := range equation.numbers[1:] {
				switch operators[j] {
				case 0:
					res += number
				case 1:
					res *= number
				case 2:
					res = aocmath.Concatenate(res, number)
				}
			}
			if res == equation.result {
				sum += equation.result
				break
			}
		}
	}

	return sum, nil
}

func parse(input string) (result []equation) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		equation := equation{}
		colon := strings.Index(line, ":")
		equation.result = aocstrconv.MustAtoi(line[:colon])

		parts := strings.Split(line[colon+2:], " ")

		for _, part := range parts {
			equation.numbers = append(equation.numbers, aocstrconv.MustAtoi(part))
		}

		result = append(result, equation)
	}

	return result
}
