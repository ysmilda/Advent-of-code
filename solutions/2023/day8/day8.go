package aoc2023day8

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	instructions string
	directions   map[string]instruction
}

type instruction struct {
	left  string
	right string
}

func (i instruction) Follow(direction byte) string {
	switch direction {
	case 'L':
		return i.left
	case 'R':
		return i.right
	}
	panic("invalid direction")
}

func MustGetSolver() solver.Solver {
	instructions, mapping := parse(inputFile)
	return puzzle{
		instructions: instructions,
		directions:   mapping,
	}
}

func (s puzzle) GetDay() int {
	return 8
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	stepIndex := 0
	node := "AAA"

	for {
		node = s.directions[node].Follow(s.instructions[stepIndex])
		if node == "ZZZ" {
			break
		}

		stepIndex++
		if stepIndex >= len(s.instructions) {
			stepIndex = 0
		}
		sum++
	}

	return sum + 1, nil
}

func (s puzzle) Part2() (int, error) {
	nodes := []string{}
	for node := range s.directions {
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	counts := make(chan int)
	for _, node := range nodes {
		node := node

		go func() {
			count := 0
			step := 0
			for {
				node = s.directions[node].Follow(s.instructions[step])
				count++
				if node[2] == 'Z' {
					break
				}
				step++
				if step >= len(s.instructions) {
					step = 0
				}
			}
			counts <- count
		}()
	}

	nodeCounts := []int{}
	for count := range counts {
		nodeCounts = append(nodeCounts, count)
		if len(nodeCounts) == len(nodes) {
			break
		}
	}

	return aocmath.LCMs(nodeCounts...), nil
}

func parse(input string) (string, map[string]instruction) {
	lines := strings.Split(input, "\n")
	instructions := lines[0]

	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	mapping := make(map[string]instruction)
	for _, line := range lines[2:] {
		line := replacer.Replace(line)
		parts := strings.Fields(line)
		mapping[parts[0]] = instruction{
			left:  parts[1],
			right: parts[2],
		}
	}

	return instructions, mapping
}
