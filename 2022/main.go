package main

import (
	"flag"
	"time"

	"advent-of-code-2022/day_1"
	"advent-of-code-2022/day_2"
	"advent-of-code-2022/day_3"
	"advent-of-code-2022/day_4"
	"advent-of-code-2022/day_5"
	"advent-of-code-2022/day_6"
	"advent-of-code-2022/day_7"

	solver "advent-of-code-2022/SolverInterface"

	log "github.com/sirupsen/logrus"
)

var solvers []solver.SolverInterface = []solver.SolverInterface{
	day_1.Solver{},
	day_2.Solver{},
	day_3.Solver{},
	day_4.Solver{},
	day_5.Solver{},
	day_6.Solver{},
	day_7.Solver{},
}

func main() {
	solverFlag := flag.Uint("s", 0, "the solver to run, when set to 0 all solvers will be run")
	flag.Parse()

	// Normalize flag input
	solverIndex := int(*solverFlag)

	// Program logic
	// Either do nothing because the requested day is not created yet
	// Or execute all solvers if the requested day == 0
	// Or execute the requested day
	switch {

	case solverIndex > len(solvers):
		log.WithFields(log.Fields{
			"solver": solverIndex,
		}).Fatal("Solution not created yet")

	case solverIndex == 0:
		for _, solver := range solvers {
			executeSolver(solver)
		}

	default:
		executeSolver(solvers[solverIndex-1])
	}
}

func executeSolver(solver solver.SolverInterface) {
	start := time.Now()

	input, err := solver.GetSolver()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Fatal("Failed to parse input")
	}

	part1, err := input.Part1()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Info("Failed to solve part 1")
	}

	part2, err := input.Part2()
	if err != nil {
		log.WithFields(log.Fields{
			"day":   solver.GetDay(),
			"error": err,
		}).Info("Failed to solve part 2")
	}

	log.WithFields(log.Fields{
		"day":     solver.GetDay(),
		"part 1":  part1,
		"part 2":  part2,
		"runtime": time.Since(start).String(),
	}).Info("Solution found")
}
