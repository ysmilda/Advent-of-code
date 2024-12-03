package aoc2023day2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []game
}

type game struct {
	id    int
	hands []hand
}

type hand struct {
	red, green, blue int
}

func (h hand) isValid(available hand) bool {
	return h.red <= available.red && h.green <= available.green && h.blue <= available.blue
}

func (h hand) multiply() int {
	return h.red * h.green * h.blue
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 2
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	available := hand{12, 13, 14}

	for _, game := range s.input {
		for _, hand := range game.hands {
			if !hand.isValid(available) {
				goto skip
			}
		}
		sum += game.id
	skip:
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, game := range s.input {
		minimum := hand{0, 0, 0}
		for _, hand := range game.hands {
			minimum.red = aocmath.Max(hand.red, minimum.red)
			minimum.green = aocmath.Max(hand.green, minimum.green)
			minimum.blue = aocmath.Max(hand.blue, minimum.blue)
		}
		sum += minimum.multiply()
	}

	return sum, nil
}

func parse(input string) []game {
	lines := strings.Split(input, "\n")

	games := []game{}
	for _, line := range lines {
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(input string) game {
	colonSplit := strings.Split(input, ":")
	id, err := strconv.Atoi(strings.Fields(colonSplit[0])[1])
	if err != nil {
		panic(err)
	}

	hands := []hand{}
	handsString := strings.Split(colonSplit[1], ";")
	for _, handString := range handsString {
		hands = append(hands, parseHand(handString))
	}

	return game{
		id:    id,
		hands: hands,
	}
}

func parseHand(input string) hand {
	colors := strings.Split(input, ",")
	var red, green, blue int
	var err error

	for _, color := range colors {
		colorSplit := strings.Fields(color)

		switch colorSplit[1] {
		case "red":
			red, err = strconv.Atoi(colorSplit[0])
		case "green":
			green, err = strconv.Atoi(colorSplit[0])
		case "blue":
			blue, err = strconv.Atoi(colorSplit[0])
		}
		if err != nil {
			panic(err)
		}
	}
	return hand{red, green, blue}
}
