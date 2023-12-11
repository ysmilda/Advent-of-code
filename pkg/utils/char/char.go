package char

import "strconv"

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func DigitToInt(c byte) int {
	return int(c - '0')
}

func MustToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}
