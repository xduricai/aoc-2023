package util

const runeToDigitOffset = 48
const hexOffsetLower = 87
const hexOffsetUpper = 55

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
	if (*input)[0] == '-' {
		return -sum
	}
	return sum
}

func ParseIntFromHexString(input *string) int {
	var sum int
	multiplier := 1

	for i := range *input {
		char := (*input)[len(*input)-i-1]

		if char >= '1' && char <= '9' {
			sum += (int(char) - runeToDigitOffset) * multiplier
		}
		if char >= 'a' && char <= 'f' {
			sum += (int(char) - hexOffsetLower) * multiplier
		}
		if char >= 'A' && char <= 'F' {
			sum += (int(char) - hexOffsetUpper) * multiplier
		}
		multiplier *= 16
	}
	if (*input)[0] == '-' {
		return -sum
	}
	return sum
}

func ParseIntsFromString(input *string) []int {
	numbers := []int{}
	strings := ParseNumbersFromString(input)

	for _, num := range strings {
		numbers = append(numbers, ParseIntFromString(&num))
	}
	return numbers
}

func ParseNumbersFromString(input *string) []string {
	numbers := []string{}

	numStart := -1
	numEnd := -1

	for idx, char := range *input {
		if IsDigit(char) || char == '-' {
			if numStart < 0 {
				numStart = idx
			}
			numEnd = idx + 1
			continue
		}
		if numStart == -1 {
			continue
		}

		numbers = append(numbers, (*input)[numStart:numEnd])
		numStart = -1
		numEnd = -1
	}
	if numStart > -1 {
		numbers = append(numbers, (*input)[numStart:numEnd])
	}

	return numbers
}

func IndexOfRune(input *string, target rune) int {
	for i := range *input {
		if rune((*input)[i]) == target {
			return i
		}
	}
	return -1
}

func Transpose(input *[]string) []string {
	width := len((*input)[0])
	height := len(*input)
	output := make([]string, width)

	for col := 0; col < width; col++ {
		line := make([]byte, height)

		for row := 0; row < height; row++ {
			line[row] = (*input)[row][col]
		}
		output[col] = string(line)
	}

	return output
}
