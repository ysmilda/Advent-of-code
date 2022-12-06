package day_5

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	solver "advent-of-code-2022/SolverInterface"
)

//go:embed input.txt
var inputFile string

type Solver struct {
	input warehouse
}

type warehouse struct {
	stacks []stack

	instructions []instruction
}

func (w warehouse) move() {
	for _, instruction := range w.instructions {
		instruction.execute(&w.stacks[instruction.source], &w.stacks[instruction.destination])
	}
}

func (w warehouse) String() string {
	output := "\r\n"
	for i, stack := range w.stacks {
		output += fmt.Sprintf("%d %s\r\n", i+1, string(stack.crates))
	}
	return output
}

type stack struct {
	crates []rune
}

// Remove removes the last entry from the stack and returns the removed entry.
func (s *stack) remove() string {
	if len(s.crates) < 1 {
		return ""
	}
	lastEntry := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return string(lastEntry)
}

func (s *stack) add(input string) {
	s.crates = append(s.crates, rune(input[0]))
}

func (s stack) last() string {
	if len(s.crates) < 1 {
		return ""
	}
	return string(s.crates[len(s.crates)-1])
}

type instruction struct {
	source      int
	destination int
	amount      int
}

func (in instruction) execute(source, destination *stack) {
	for i := 0; i < in.amount; i++ {
		destination.add(source.remove())
	}
}

func (in instruction) executeInOrder(source, destination *stack) {
	buffer := []rune{}
	for i := 0; i < in.amount; i++ {
		buffer = append(buffer, rune(source.remove()[0]))
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		destination.add(string(buffer[i]))
	}
}

func (s Solver) GetSolver() (solver.SolverInterface, error) {
	input, err := parseInput()
	if err != nil {
		return Solver{}, err
	}

	return Solver{input: input}, nil
}

func (s Solver) GetDay() int {
	return 5
}

func (s Solver) Part1() (string, error) {
	output := ""
	stacks := make([]stack, len(s.input.stacks))
	copy(stacks, s.input.stacks)

	for _, instruction := range s.input.instructions {
		instruction.execute(&stacks[instruction.source], &stacks[instruction.destination])
	}

	for _, stack := range stacks {
		output += string(stack.last())
	}

	return output, nil
}

func (s Solver) Part2() (string, error) {
	output := ""
	stacks := make([]stack, len(s.input.stacks))
	copy(stacks, s.input.stacks)

	for _, instruction := range s.input.instructions {
		instruction.executeInOrder(&stacks[instruction.source], &stacks[instruction.destination])
	}

	for _, stack := range stacks {
		output += string(stack.last())
	}

	return output, nil
}

func parseInput() (warehouse, error) {
	lines := strings.Split(inputFile, "\n")

	currentStacks := []string{}
	output := warehouse{
		stacks: make([]stack, 10),
	}
	instructionsInput := false
	for _, line := range lines {
		if line == "" {
			// Now that we have the complete warehouse layout, let's parse it.
			for i := 1; i < 10; i++ {
				index := strings.Index(currentStacks[len(currentStacks)-1], strconv.Itoa(i))
				if index == -1 {
					continue
				}
				stack := stack{}
				for j := len(currentStacks) - 2; j >= 0; j-- {
					if currentStacks[j][index] == 32 {
						continue
					}
					stack.crates = append(stack.crates, rune(currentStacks[j][index]))
				}
				output.stacks[i-1] = stack
			}

			// And switch over to reading the instructions
			instructionsInput = true
			continue
		}

		switch instructionsInput {
		case false: // Parse current stack situation
			currentStacks = append(currentStacks, line)

		case true: // Parse the line as a instruction
			splitLine := strings.Split(line, " ")
			destination, _ := strconv.Atoi(splitLine[5])
			source, _ := strconv.Atoi(splitLine[3])
			amount, _ := strconv.Atoi(splitLine[1])
			output.instructions = append(output.instructions, instruction{
				amount:      amount,
				destination: destination - 1,
				source:      source - 1,
			})
		}
	}

	return output, nil
}
