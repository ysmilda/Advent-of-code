package utils

func CharIsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func CharToInt(char byte) int {
	if !CharIsDigit(char) {
		return 0
	}
	return int(char) - '0'
}
