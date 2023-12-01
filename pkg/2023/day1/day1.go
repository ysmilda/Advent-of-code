package aoc2023day1

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
	"github.com/ysmilda/Advent-of-code/pkg/utils"
)

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
	return 1
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	for _, line := range s.input {
		digits := findDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")

	for _, line := range s.input {
		line := replacer.Replace(line)
		digits := findDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum, nil
}

func findDigits(line string) []int {
	digits := []int{}
	for i := 0; i < len(line); i++ {
		char := line[i]
		if utils.CharIsDigit(char) {
			digits = append(digits, utils.CharToInt(char))
		}
	}
	return digits
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}
