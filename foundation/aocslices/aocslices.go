package aocslices

import (
	"iter"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
)

// Remove removes the element at index from a slice.
func Remove[T []E, E any](slice T, index int) T {
	return append(slice[:index], slice[index+1:]...)
}

// Iterate returns an iterater which iterates a number in a given base and returns each digit as part of a slice.
func Iterate(base, length int) iter.Seq2[int, []int] {
	return func(yield func(int, []int) bool) {
		steps := make([]int, length)
		for i := range aocmath.Pow(base, length) {
			if !yield(i, steps) {
				return
			}
			for j := range steps {
				steps[j]++
				if steps[j] > base-1 {
					steps[j] = 0
				} else {
					break
				}
			}
		}
	}
}
