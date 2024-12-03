package aoc2024day3

import (
	_ "embed"
	"regexp"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

var multiplication = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
var all = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []string
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 3
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, instruction := range s.input {
		matches := multiplication.FindAll([]byte(instruction), -1)

		for _, match := range matches {
			sum += parseAndMultiply(string(match[4 : len(match)-1]))
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	enabled := true
	for _, instruction := range s.input {
		matches := all.FindAll([]byte(instruction), -1)

		for _, match := range matches {
			switch string(match) {
			case "do()":
				enabled = true

			case "don't()":
				enabled = false

			default:
				if enabled {
					sum += parseAndMultiply(string(match[4 : len(match)-1]))
				}
			}
		}
	}

	return sum, nil
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func parseAndMultiply(s string) int {
	parts := strings.Split(s, ",")
	return (aocstrconv.MustAtoi(parts[0]) * aocstrconv.MustAtoi(parts[1]))
}
