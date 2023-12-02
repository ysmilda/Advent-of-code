package aoc2023day2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
	"github.com/ysmilda/Advent-of-code/pkg/utils"
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
			minimum.red = utils.Max(hand.red, minimum.red)
			minimum.green = utils.Max(hand.green, minimum.green)
			minimum.blue = utils.Max(hand.blue, minimum.blue)
		}
		sum += minimum.multiply()
	}

	return sum, nil
}

func parse(input string) []game {
	lines := strings.Split(input, "\n")

	games := []game{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(input string) game {
	colonSplit := split(input, ":")
	id, err := strconv.Atoi(split(colonSplit[0], " ")[1])
	if err != nil {
		panic(err)
	}

	hands := []hand{}
	handsString := split(colonSplit[1], ";")
	for _, handString := range handsString {
		hands = append(hands, parseHand(handString))
	}

	return game{
		id:    id,
		hands: hands,
	}
}

func parseHand(input string) hand {
	colors := split(input, ",")
	var red, green, blue int
	var err error

	for _, color := range colors {
		colorSplit := split(color, " ")

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

func split(s string, sep string) []string {
	entries := strings.Split(s, sep)

	for i := range entries {
		entries[i] = strings.TrimSpace(entries[i])
	}
	return entries
}
