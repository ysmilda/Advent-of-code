package utils

import "golang.org/x/exp/constraints"

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

func Abs[T constraints.Integer | constraints.Float](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func DigitToInt(c byte) int {
	return int(c - '0')
}
