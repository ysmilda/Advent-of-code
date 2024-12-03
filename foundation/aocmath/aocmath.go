package aocmath

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Limit[T constraints.Integer | constraints.Float](value, min, max T) T {
	return Min(Max(value, min), max)
}

func Between[T constraints.Integer | constraints.Float](value, a, b T) bool {
	return value > Min(a, b) && value < Max(a, b)
}

func BetweenInclusive[T constraints.Integer | constraints.Float](value, a, b T) bool {
	return value >= Min(a, b) && value <= Max(a, b)
}

func Abs[T constraints.Integer | constraints.Float](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

func SameSign[T constraints.Integer | constraints.Float](a, b T) bool {
	return a <= 0 && b <= 0 || a >= 0 && b >= 0
}

func GCD[T constraints.Integer](a, b T) T {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

func GCDs[T constraints.Integer](values ...T) T {
	result := values[0]
	for _, value := range values[1:] {
		result = GCD(result, value)
	}
	return result
}

func LCM[T constraints.Integer](a, b T) T {
	if a == 0 || b == 0 {
		panic("LCM of 0 is undefined")
	}
	return a * b / GCD(a, b)
}

func LCMs[T constraints.Integer](values ...T) T {
	result := values[0]
	for _, value := range values[1:] {
		result = LCM(result, value)
	}
	return result
}
