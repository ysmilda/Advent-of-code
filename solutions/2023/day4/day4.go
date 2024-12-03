package aoc2023day4

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	cards []card
}

type card struct {
	instances       int
	matchingNumbers int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		cards: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 4
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, card := range s.cards {
		i := card.matchingNumbers
		if i != 0 {
			sum += (1 << (i - 1))
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for i := 0; i < len(s.cards); i++ {
		for j := 1; j <= s.cards[i].matchingNumbers; j++ {
			s.cards[i+j].instances += s.cards[i].instances
		}
		sum += s.cards[i].instances
	}

	return sum, nil
}

func parse(input string) []card {
	lines := strings.Split(input, "\n")

	output := []card{}
	for _, line := range lines {
		numbersSplit := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := convertToNumbers(strings.Fields(numbersSplit[0]))
		entryNumbers := convertToNumbers(strings.Fields(numbersSplit[1]))

		output = append(output, card{
			instances:       1,
			matchingNumbers: matchingNumbers(winningNumbers, entryNumbers),
		})
	}

	return output
}

func matchingNumbers(winningNumbers, numbers []int) (output int) {
	for _, winningNumber := range winningNumbers {
		for _, number := range numbers {
			if winningNumber == number {
				output++
			}
		}
	}

	return output
}

func convertToNumbers(input []string) []int {
	output := []int{}

	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		output = append(output, num)
	}

	return output
}
