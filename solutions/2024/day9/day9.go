package aoc2024day9

import (
	_ "embed"
	"slices"

	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 9
}

func (s puzzle) Part1() (int, error) {
	disk := make([]int, len(s.input))
	copy(disk, s.input)

	// Fill the empty space with the last file value from the disk.
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			last := disk[len(disk)-1]

			for last == -1 {
				disk = disk[:len(disk)-1]
				last = disk[len(disk)-1]
			}

			disk[i] = last
			disk = disk[:len(disk)-1]
		}
	}

	return checksum(disk), nil
}

func (s puzzle) Part2() (int, error) {
	disk := make([]int, len(s.input))
	copy(disk, s.input)

	// Find the first empty index so we can use that to exit early below.
	firstEmpty := 0
	for i := range len(disk) {
		if disk[i] == -1 {
			firstEmpty = i
			break
		}
	}

	// Loop backwards over the disk, find the length of a file, and try and put it forward as much as possible.
	for fileStart := len(disk) - 1; fileStart > firstEmpty; fileStart-- {
		if disk[fileStart] == -1 {
			continue
		}

		// Find the file fileLength
		fileLength := 0
		for fileEnd := fileStart; fileEnd > 0; fileEnd-- {
			if disk[fileEnd] != disk[fileStart] {
				fileLength = fileStart - fileEnd
				break
			}
		}

		// Look for empty space in front of file.
		for emptyStart := 0; emptyStart < fileStart; emptyStart++ {
			if disk[emptyStart] != -1 {
				continue
			}

			fits := false
			emptyEnd := emptyStart
			for ; emptyEnd < len(disk); emptyEnd++ {
				if disk[emptyEnd] != -1 {
					break
				}
				// We found a space where it will fit.
				if (emptyEnd-emptyStart)+1 == fileLength {
					fits = true
					break
				}
			}

			// Copy the data from the file to the empty space.
			if fits {
				fileContent := disk[fileStart]
				for l := range fileLength {
					disk[emptyStart+l] = fileContent
					disk[fileStart-l] = -1
				}
				break
			}

			emptyStart = emptyEnd - 1
		}
		fileStart -= fileLength - 1
	}

	return checksum(disk), nil
}

func parse(input string) []int {
	result := []int{}
	for i, c := range input {
		if c == '0' {
			continue
		}
		if i%2 == 0 {
			result = append(result, slices.Repeat([]int{i / 2}, char.ToInt(byte(c)))...)
		} else {
			result = append(result, slices.Repeat([]int{-1}, char.ToInt(byte(c)))...)
		}
	}
	return result
}

func checksum(in []int) int {
	checksum := 0
	for i, value := range in {
		if value == -1 {
			continue
		}
		checksum += i * value
	}
	return checksum
}
