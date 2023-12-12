package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCoordinate(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected Coordinate
	}{
		{0, 0, Coordinate{0, 0}},
		{1, 2, Coordinate{1, 2}},
		{3, 4, Coordinate{3, 4}},
	}

	for _, tc := range testCases {
		result := NewCoordinate(tc.x, tc.y)
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateAdd(t *testing.T) {
	testCases := []struct {
		c1, c2, expected Coordinate
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}, Coordinate{0, 0}},
		{Coordinate{1, 2}, Coordinate{3, 4}, Coordinate{4, 6}},
		{Coordinate{5, 6}, Coordinate{7, 8}, Coordinate{12, 14}},
	}

	for _, tc := range testCases {
		result := tc.c1.Add(tc.c2)
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateAddDirection(t *testing.T) {
	testCases := []struct {
		c         Coordinate
		d         Direction
		increment uint
		expected  Coordinate
	}{
		{Coordinate{0, 0}, North, 1, Coordinate{0, -1}},
		{Coordinate{0, 0}, NorthEast, 1, Coordinate{1, -1}},
		{Coordinate{0, 0}, East, 1, Coordinate{1, 0}},
		{Coordinate{0, 0}, SouthEast, 1, Coordinate{1, 1}},
		{Coordinate{0, 0}, South, 1, Coordinate{0, 1}},
		{Coordinate{0, 0}, SouthWest, 1, Coordinate{-1, 1}},
		{Coordinate{0, 0}, West, 1, Coordinate{-1, 0}},
		{Coordinate{0, 0}, NorthWest, 1, Coordinate{-1, -1}},
		{Coordinate{0, 0}, None, 1, Coordinate{0, 0}},
	}

	for _, tc := range testCases {
		result := tc.c.AddDirection(tc.d, tc.increment)
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateSubtract(t *testing.T) {
	testCases := []struct {
		c1, c2, expected Coordinate
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}, Coordinate{0, 0}},
		{Coordinate{1, 2}, Coordinate{3, 4}, Coordinate{-2, -2}},
		{Coordinate{5, 6}, Coordinate{7, 8}, Coordinate{-2, -2}},
	}

	for _, tc := range testCases {
		result := tc.c1.Subtract(tc.c2)
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateMultiply(t *testing.T) {
	testCases := []struct {
		c1, c2, expected Coordinate
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}, Coordinate{0, 0}},
		{Coordinate{1, 2}, Coordinate{3, 4}, Coordinate{3, 8}},
		{Coordinate{5, 6}, Coordinate{7, 8}, Coordinate{35, 48}},
	}

	for _, tc := range testCases {
		result := tc.c1.Multiply(tc.c2)
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateDivide(t *testing.T) {
	testCases := []struct {
		c1, c2, expected Coordinate
	}{
		{Coordinate{2, 4}, Coordinate{2, 2}, Coordinate{1, 2}},
		{Coordinate{15, 25}, Coordinate{2, 2}, Coordinate{7, 12}},
	}

	for _, tc := range testCases {
		result := tc.c1.Divide(tc.c2)
		assert.Equal(t, tc.expected, result)
	}

	assert.Panics(t, func() { Coordinate{0, 0}.Divide(Coordinate{0, 0}) })
}

func TestCoordinateAbs(t *testing.T) {
	testCases := []struct {
		c        Coordinate
		expected Coordinate
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}},
		{Coordinate{1, 2}, Coordinate{1, 2}},
		{Coordinate{-1, -2}, Coordinate{1, 2}},
	}

	for _, tc := range testCases {
		result := tc.c.Abs()
		assert.Equal(t, tc.expected, result)
	}
}

func TestCoordinateManhattanDistance(t *testing.T) {
	testCases := []struct {
		c1, c2   Coordinate
		expected int
	}{
		{Coordinate{0, 0}, Coordinate{0, 0}, 0},
		{Coordinate{1, 2}, Coordinate{3, 4}, 4},
		{Coordinate{5, 6}, Coordinate{7, 8}, 4},
	}

	for _, tc := range testCases {
		result := tc.c1.ManhattanDistance(tc.c2)
		assert.Equal(t, tc.expected, result)
	}
}
