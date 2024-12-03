package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

var solvers map[int][]solver.Solver

func main() {
	year := flag.Int("year", 2024, "the year to run")
	day := flag.Int("day", 0, "the solver to run, when set to 0 all solvers will be run")
	flag.Parse()

	s, ok := solvers[*year]
	if !ok {
		panic("Year not implemented yet")
	}

	ran := true
	for _, solver := range s {
		if solver.GetDay() == *day || *day == 0 {
			executeSolver(solver)
		}
	}

	if !ran {
		fmt.Println("Day not implemented yet")
	}
}

func executeSolver(solver solver.Solver) {
	printDay(solver.GetDay())

	start := time.Now()
	solution, err := solver.Part1()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}

	dur := time.Since(start)
	printPart(1, solution, dur)

	solution, err = solver.Part2()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}

	dur = time.Since(start) - dur
	printPart(2, solution, dur)
}

func printDay(day int) {
	fmt.Printf("Day %d:\n", day)
}

func printPart(part int, solution int, duration time.Duration) {
	fmt.Printf("\tPart %d:\n\t\tsolution: %d\n\t\truntime: %s\n", part, solution, duration.String())
}
