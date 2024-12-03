package aoc2023day3

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	symbols     map[symbol][]int
	partNumbers []int
}

type symbol struct {
	symbol byte
	x      int
	y      int
}

func MustGetSolver() solver.Solver {
	partNumbers, symbols := parse(inputFile)
	return puzzle{
		partNumbers: partNumbers,
		symbols:     symbols,
	}
}

func (s puzzle) GetDay() int {
	return 3
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, partNumber := range s.partNumbers {
		sum += partNumber
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for symbol, partNumbers := range s.symbols {
		if symbol.symbol == '*' {
			if len(partNumbers) < 2 {
				continue
			}

			ratio := 1
			for _, partNumber := range partNumbers {
				ratio *= partNumber
			}
			sum += ratio
		}
	}

	return sum, nil
}

func parse(input string) ([]int, map[symbol][]int) {
	lines := strings.Split(input, "\n")

	symbols := make(map[symbol][]int)
	partNumbers := []int{}

	for i := 0; i < len(lines); i++ {
		foundPartNumber := false
		partNumber := 0
		foundSymbols := make(map[symbol]any)

		for j := 0; j < len(lines[i]); j++ {
			if char.IsDigit(lines[i][j]) {
				partNumber = partNumber*10 + char.ToInt(lines[i][j])

				// Check for symbols around the number
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if k < 0 || l < 0 || k > len(lines)-1 || l > len(lines[k])-1 {
							continue
						}

						if lines[k][l] != '.' && !char.IsDigit(lines[k][l]) {
							foundSymbols[symbol{
								symbol: lines[k][l],
								x:      k,
								y:      l,
							}] = nil

							foundPartNumber = true
						}
					}
				}
			} else {
				if foundPartNumber {
					partNumbers = append(partNumbers, partNumber)

					for symbol := range foundSymbols {
						symbols[symbol] = append(symbols[symbol], partNumber)
					}
				}

				foundPartNumber = false
				partNumber = 0
				foundSymbols = make(map[symbol]any)
			}
		}
	}

	return partNumbers, symbols
}
