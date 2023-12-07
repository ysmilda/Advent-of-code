package aoc2023day7

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
	"github.com/ysmilda/Advent-of-code/pkg/utils"
)

type cardStrength uint

const (
	highCard cardStrength = iota + 1
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	strength cardStrength
	value    int
	bid      int
}

//go:embed input.txt
var inputFile string

type puzzle struct {
	input string
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: inputFile,
	}
}

func (s puzzle) GetDay() int {
	return 7
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	input := sort(parse(s.input, false, "23456789TJQKA"))

	for i, in := range input {
		sum += (in.bid * (i + 1))
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0
	input := sort(parse(s.input, true, "J23456789TQKA"))
	for i, in := range input {
		sum += (in.bid * (i + 1))
	}

	return sum, nil
}

func parse(input string, useJokers bool, values string) []hand {
	lines := strings.Split(input, "\n")

	output := []hand{}
	for _, line := range lines {
		output = append(output, parseHand(line, useJokers, values))
	}
	return output
}

func parseHand(line string, useJokers bool, values string) hand {
	parts := strings.Fields(line)

	cards := make(map[byte]int)
	value := 0
	for _, c := range parts[0] {
		cards[byte(c)]++
		value = (value + strings.IndexByte(values, byte(c))) * 100
	}

	jokers := cards['J']
	if useJokers {
		delete(cards, 'J')
	}
	cardStrength := highCard
	pairs := 0
	threeOfAKindFound := false
	for _, v := range cards {
		if useJokers {
			v += jokers
		}
		switch v {
		case 5:
			cardStrength = fiveOfAKind
			break
		case 4:
			cardStrength = fourOfAKind
			break
		case 3:
			threeOfAKindFound = true
		case 2:
			pairs++
		}
	}
	if cardStrength == highCard {
		if threeOfAKindFound && pairs == 1 {
			cardStrength = fullHouse
		} else if threeOfAKindFound {
			cardStrength = threeOfAKind
		} else if pairs == 2 {
			cardStrength = twoPair
		} else if pairs == 1 {
			cardStrength = onePair
		}
	}

	return hand{
		bid:      utils.MustToInt(parts[1]),
		strength: cardStrength,
		value:    value,
	}
}

func sort(input []hand) []hand {
	slices.SortFunc(input, func(a, b hand) int {
		if a.strength == b.strength {
			if a.value > b.value {
				return +1
			}
			if a.value < b.value {
				return -1
			}
			return 0
		}
		if a.strength > b.strength {
			return +1
		}
		return -1
	})
	return input
}
