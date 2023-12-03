package util

const runeToDigitOffset = 48

func IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func IsBDigit(b byte) bool {
	char := rune(b)
	return char >= '0' && char <= '9'
}

func ParseIntFromRune(b byte) int {
	r := rune(b)
	if r < '1' || r > '9' {
		return 0
	}

	return int(r) - runeToDigitOffset
}

func ParseIntFromString(input *string) int {
	var sum int
	multiplier := 1

	for i := range *input {
		char := (*input)[len(*input)-i-1]

		if char >= '1' && char <= '9' {
			sum += (int(char) - runeToDigitOffset) * multiplier
		}
		multiplier *= 10
	}
	return sum
}

func IndexOfRune(input *string, target rune) int {
	for i := range *input {
		if rune((*input)[i]) == target {
			return i
		}
	}
	return -1
}
