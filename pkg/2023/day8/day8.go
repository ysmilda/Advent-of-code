package aoc2023day8

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
	"github.com/ysmilda/Advent-of-code/pkg/utils"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	instructions string
	mapping      map[string]instruction
}

type instruction struct {
	left  string
	right string
}

func MustGetSolver() solver.Solver {
	instructions, mapping := parse(inputFile)
	return puzzle{
		instructions: instructions,
		mapping:      mapping,
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
		switch s.instructions[stepIndex] {
		case 'L':
			node = s.mapping[node].left
		case 'R':
			node = s.mapping[node].right
		}
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
	stepIndex := 0

	nodes := []string{}
	for node := range s.mapping {
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	nodeCounts := make([]int, len(nodes))
	finishedNodes := make([]bool, len(nodes))
	for {
		for i := 0; i < len(nodes); i++ {
			if finishedNodes[i] {
				continue
			}
			switch s.instructions[stepIndex] {
			case 'L':
				nodes[i] = s.mapping[nodes[i]].left
			case 'R':
				nodes[i] = s.mapping[nodes[i]].right
			}
			nodeCounts[i]++
			if nodes[i][2] == 'Z' {
				finishedNodes[i] = true
			}
		}

		allFinished := true
		for _, finished := range finishedNodes {
			if !finished {
				allFinished = false
				break
			}
		}
		if allFinished {
			break
		}

		stepIndex++
		if stepIndex >= len(s.instructions) {
			stepIndex = 0
		}
	}

	return utils.LCMs(nodeCounts...), nil
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
