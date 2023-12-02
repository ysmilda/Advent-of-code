package aoc2023day2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
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
			if hand.red > available.red || hand.green > available.green || hand.blue > available.blue {
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
			if hand.red > minimum.red {
				minimum.red = hand.red
			}
			if hand.green > minimum.green {
				minimum.green = hand.green
			}
			if hand.blue > minimum.blue {
				minimum.blue = hand.blue
			}
		}
		sum += (minimum.red * minimum.blue * minimum.green)
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

		gameSplit := split(line, ":")
		id, err := strconv.Atoi(split(gameSplit[0], " ")[1])
		if err != nil {
			panic(err)
		}

		var red, green, blue int
		hands := []hand{}
		for _, entry := range split(gameSplit[1], ";") {
			colors := split(entry, ",")

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
			hands = append(hands, hand{red, green, blue})
		}
		games = append(games, game{
			id:    id,
			hands: hands,
		})
	}
	return games
}

func split(s string, sep string) []string {
	entries := strings.Split(s, sep)

	for i := range entries {
		entries[i] = strings.TrimSpace(entries[i])
	}
	return entries
}
