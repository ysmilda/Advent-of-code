package main

import (
	"flag"
	"fmt"
	"time"

	aoc2023day1 "github.com/ysmilda/Advent-of-code/pkg/2023/day1"
	aoc2023day2 "github.com/ysmilda/Advent-of-code/pkg/2023/day2"
	aoc2023day3 "github.com/ysmilda/Advent-of-code/pkg/2023/day3"
	aoc2023day4 "github.com/ysmilda/Advent-of-code/pkg/2023/day4"
	aoc2023day5 "github.com/ysmilda/Advent-of-code/pkg/2023/day5"
	aoc2023day6 "github.com/ysmilda/Advent-of-code/pkg/2023/day6"
	aoc2023day7 "github.com/ysmilda/Advent-of-code/pkg/2023/day7"
	aoc2023day8 "github.com/ysmilda/Advent-of-code/pkg/2023/day8"

	"github.com/ysmilda/Advent-of-code/pkg/solver"
)

var solvers map[int]map[int]solver.Solver

func init() {
	start := time.Now()
	solvers = map[int]map[int]solver.Solver{
		2023: {
			1: aoc2023day1.MustGetSolver(),
			2: aoc2023day2.MustGetSolver(),
			3: aoc2023day3.MustGetSolver(),
			4: aoc2023day4.MustGetSolver(),
			5: aoc2023day5.MustGetSolver(),
			6: aoc2023day6.MustGetSolver(),
			7: aoc2023day7.MustGetSolver(),
			8: aoc2023day8.MustGetSolver(),
		},
	}
	printParseTime(time.Since(start))
}

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
