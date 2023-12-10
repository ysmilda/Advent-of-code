package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
)

var solvers map[int]map[int]solver.Solver

func main() {
	year := flag.Int("year", 2023, "the year to run")
	day := flag.Int("day", 0, "the solver to run, when set to 0 all solvers will be run")
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

func printParseTime(dur time.Duration) {
	fmt.Printf("parse time: %s\n", dur.String())
}

func printDay(day int) {
	fmt.Printf("Day %d:\n", day)
}

func printPart(part int, solution int, duration time.Duration) {
	fmt.Printf("\tPart %d:\n\t\tsolution: %d\n\t\truntime: %s\n", part, solution, duration.String())
}
