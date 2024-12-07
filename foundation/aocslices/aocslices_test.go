package aocslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	expectedResults := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	}
	j := 0
	for _, i := range Iterate(2, 2) {
		assert.Equal(t, expectedResults[j], i)
		j++
	}
	assert.Equal(t, len(expectedResults), j, "Not all results were iterated over")

	expectedResults = [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}
	j = 0
	for _, i := range Iterate(3, 2) {
		assert.Equal(t, expectedResults[j], i)
		j++
	}
	assert.Equal(t, len(expectedResults), j, "Not all results were iterated over")

}
