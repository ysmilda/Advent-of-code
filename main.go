package main

import (
	"flag"
	"fmt"
	"time"

	aoc2023 "github.com/ysmilda/Advent-of-code/pkg/2023"
	"github.com/ysmilda/Advent-of-code/pkg/solver"
)

var solvers = map[uint]map[uint]solver.Solver{
	2023: {
		1: aoc2023.MustGetSolver2023Day1(),
	},
}

func main() {
	year := flag.Uint("year", 2023, "the year to run, when set to 0 all years will be run")
	day := flag.Uint("day", 0, "the solver to run, when set to 0 all solvers will be run")
	flag.Parse()

	s, ok := solvers[*year]
	if !ok {
		panic("Year not implemented yet")
	}

	if *day == 0 {
		for _, solver := range s {
			executeSolver(solver)
		}
		return
	}

	if _, ok := s[*day]; !ok {
		panic("Day not implemented yet")
	}
	executeSolver(s[*day])
}

func executeSolver(solver solver.Solver) {
	start := time.Now()

	solutionPart1, err := solver.Part1()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}

	durationPart1 := time.Since(start)

	solutionPart2, err := solver.Part2()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}

	durationPart2 := time.Since(start) - durationPart1

	fmt.Printf("Day %d:\n\tPart 1:\n\t\tsolution: %d\n\t\truntime: %s\n\tPart 2:\n\t\tsolution: %d\n\t\truntime: %s\n", solver.GetDay(), solutionPart1, durationPart1.String(), solutionPart2, durationPart2.String())
}
