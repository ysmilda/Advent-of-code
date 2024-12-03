package aocmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 10, 5},
		{-3, 2, -3},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Min(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 10, 10},
		{-3, 2, 2},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}
}

func TestLimit(t *testing.T) {
	tests := []struct {
		value, min, max, expected int
	}{
		{5, 0, 10, 5},
		{-3, 0, 10, 0},
		{15, 0, 10, 10},
	}

	for _, test := range tests {
		result := Limit(test.value, test.min, test.max)
		assert.Equal(t, test.expected, result)
	}
}

func TestBetween(t *testing.T) {
	testCases := []struct {
		value, a, b int
		expected    bool
	}{
		{5, 0, 10, true},
		{-3, 0, 10, false},
		{15, 0, 10, false},
		{5, 10, 0, true},
		{5, 5, 10, false},
	}

	for _, tc := range testCases {
		result := Between(tc.value, tc.a, tc.b)
		assert.Equal(t, tc.expected, result)
	}
}

func TestBetweenInclusive(t *testing.T) {
	testCases := []struct {
		value, a, b int
		expected    bool
	}{
		{5, 0, 10, true},
		{-3, 0, 10, false},
		{15, 0, 10, false},
		{5, 10, 0, true},
		{5, 5, 10, true},
		{5, 0, 5, true},
	}

	for _, tc := range testCases {
		result := BetweenInclusive(tc.value, tc.a, tc.b)
		assert.Equal(t, tc.expected, result)
	}
}

func TestAbs(t *testing.T) {
	testInts := []struct {
		value, expected int
	}{
		{5, 5},
		{-3, 3},
		{0, 0},
	}

	for _, test := range testInts {
		result := Abs(test.value)
		assert.Equal(t, test.expected, result)
	}

	testFloats := []struct {
		value, expected float64
	}{
		{5.0, 5.0},
		{-3.0, 3.0},
		{0.0, 0.0},
	}

	for _, test := range testFloats {
		result := Abs(test.value)
		assert.Equal(t, test.expected, result)
	}
}

func TestSameSign(t *testing.T) {
	testInts := []struct {
		a, b     int
		expected bool
	}{
		{5, 5, true},
		{-3, 3, false},
		{0, 0, true},
	}

	for _, test := range testInts {
		result := SameSign(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}

	testFloats := []struct {
		a, b     float64
		expected bool
	}{
		{5.0, 5.0, true},
		{-3.0, 3.0, false},
		{0.0, 0.0, true},
	}

	for _, test := range testFloats {
		result := SameSign(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 10, 5},
		{3, 2, 1},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}
}

func TestGCDs(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{[]int{5, 10}, 5},
		{[]int{3, 2, 1}, 1},
		{[]int{0, 0, 0}, 0},
	}

	for _, test := range tests {
		result := GCDs(test.values...)
		assert.Equal(t, test.expected, result)
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 10, 10},
		{3, 2, 6},
	}

	for _, test := range tests {
		result := LCM(test.a, test.b)
		assert.Equal(t, test.expected, result)
	}

	assert.Panics(t, func() { LCM(0, 0) })
}

func TestLCMs(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{[]int{5, 10}, 10},
		{[]int{3, 2, 1}, 6},
	}

	for _, test := range tests {
		result := LCMs(test.values...)
		assert.Equal(t, test.expected, result)
	}
}
