package main

import (
	"github.com/ysmilda/Advent-of-code/foundation/solver"
	aoc2023day1 "github.com/ysmilda/Advent-of-code/solutions/2023/day1"
	aoc2023day10 "github.com/ysmilda/Advent-of-code/solutions/2023/day10"
	aoc2023day11 "github.com/ysmilda/Advent-of-code/solutions/2023/day11"
	aoc2023day2 "github.com/ysmilda/Advent-of-code/solutions/2023/day2"
	aoc2023day3 "github.com/ysmilda/Advent-of-code/solutions/2023/day3"
	aoc2023day4 "github.com/ysmilda/Advent-of-code/solutions/2023/day4"
	aoc2023day5 "github.com/ysmilda/Advent-of-code/solutions/2023/day5"
	aoc2023day6 "github.com/ysmilda/Advent-of-code/solutions/2023/day6"
	aoc2023day7 "github.com/ysmilda/Advent-of-code/solutions/2023/day7"
	aoc2023day8 "github.com/ysmilda/Advent-of-code/solutions/2023/day8"
	aoc2023day9 "github.com/ysmilda/Advent-of-code/solutions/2023/day9"
	aoc2024day1 "github.com/ysmilda/Advent-of-code/solutions/2024/day1"
	aoc2024day2 "github.com/ysmilda/Advent-of-code/solutions/2024/day2"
	aoc2024day3 "github.com/ysmilda/Advent-of-code/solutions/2024/day3"
	aoc2024day4 "github.com/ysmilda/Advent-of-code/solutions/2024/day4"
	aoc2024day5 "github.com/ysmilda/Advent-of-code/solutions/2024/day5"
	aoc2024day6 "github.com/ysmilda/Advent-of-code/solutions/2024/day6"
	aoc2024day7 "github.com/ysmilda/Advent-of-code/solutions/2024/day7"
)

func init() {
	solvers = map[int][]solver.Solver{
		2023: {
			aoc2023day1.MustGetSolver(),
			aoc2023day2.MustGetSolver(),
			aoc2023day3.MustGetSolver(),
			aoc2023day4.MustGetSolver(),
			aoc2023day5.MustGetSolver(),
			aoc2023day6.MustGetSolver(),
			aoc2023day7.MustGetSolver(),
			aoc2023day8.MustGetSolver(),
			aoc2023day9.MustGetSolver(),
			aoc2023day10.MustGetSolver(),
			aoc2023day11.MustGetSolver(),
		},
		2024: {
			aoc2024day1.MustGetSolver(),
			aoc2024day2.MustGetSolver(),
			aoc2024day3.MustGetSolver(),
			aoc2024day4.MustGetSolver(),
			aoc2024day5.MustGetSolver(),
			aoc2024day6.MustGetSolver(),
			aoc2024day7.MustGetSolver(),
		},
	}
}
