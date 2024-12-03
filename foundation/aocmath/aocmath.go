package aocmath

import (
	"golang.org/x/exp/constraints"
)

// Min returns the minimum of two values.
func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values.
func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Limit returns a value limited between a minimum and a maximum.
func Limit[T constraints.Integer | constraints.Float](value, min, max T) T {
	return Min(Max(value, min), max)
}

// Between returns true if a value is between two other values.
func Between[T constraints.Integer | constraints.Float](value, a, b T) bool {
	return value > Min(a, b) && value < Max(a, b)
}

// BetweenInclusive returns true if a value is between two other values, inclusive.
func BetweenInclusive[T constraints.Integer | constraints.Float](value, a, b T) bool {
	return value >= Min(a, b) && value <= Max(a, b)
}

// Abs returns the absolute value of a number.
func Abs[T constraints.Integer | constraints.Float](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

// SameSign returns true if two values have the same sign.
func SameSign[T constraints.Integer | constraints.Float](a, b T) bool {
	return a <= 0 && b <= 0 || a >= 0 && b >= 0
}

// GCD returns the greatest common divisor of two values.
func GCD[T constraints.Integer](a, b T) T {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

// GCDs returns the greatest common divisor of multiple values.
func GCDs[T constraints.Integer](values ...T) T {
	result := values[0]
	for _, value := range values[1:] {
		result = GCD(result, value)
	}
	return result
}

// LCM returns the least common multiple of two values.
func LCM[T constraints.Integer](a, b T) T {
	if a == 0 || b == 0 {
		panic("LCM of 0 is undefined")
	}
	return a * b / GCD(a, b)
}

// LCMs returns the least common multiple of multiple values.
func LCMs[T constraints.Integer](values ...T) T {
	result := values[0]
	for _, value := range values[1:] {
		result = LCM(result, value)
	}
	return result
}
