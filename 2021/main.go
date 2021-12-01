package main

import (
	day_1 "advent-of-code-2021/Day_1"
	solver "advent-of-code-2021/SolverInterface"
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
)

var solvers []solver.SolverInterface

func main() {
	dayFlag := flag.Uint("d", 0, "the day to solve, when set to 0 all days will be solved")
	flag.Parse()

	// Normalize flag input
	day := int(*dayFlag)

	// Initialize all the solvers, parsing the input.
	initSolvers()

	// Program logic
	// Either do nothing because the requested day is not created yet
	// Or execute all solvers if the requested day == 0
	// Or execute the requested day
	if day > len(solvers) {
		log.WithFields(log.Fields{
			"day": day,
		}).Fatal("Solution not created yet")

	} else if day == 0 {
		for _, solver := range solvers {
			executeSolver(solver)
		}

	} else {
		executeSolver(solvers[day-1])
	}
}

func initSolvers() {
	day_1, err := day_1.GetSolver()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"day":   1,
		}).Fatal("Failed to read input")
	}

	solvers = append(solvers,
		day_1,
	)
}

func executeSolver(solver solver.SolverInterface) {
	start := time.Now()

	log.WithFields(log.Fields{
		"day":     solver.GetDay(),
		"part 1":  solver.Part1(),
		"part 2":  solver.Part2(),
		"runtime": time.Since(start).String(),
	}).Info("Solution found")
}
