package day02

import "github.com/xduricai/aoc-2023/util"

const redMaxCount = 12
const greenMaxCount = 13
const blueMaxCount = 14
const runeToDigitOffset = 48

func SumValidGames() (int, error) {
	id := "02-1"
	games, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var sum int
	for _, line := range games {
		id := parseGameId(&line)
		if validateGame(&line) {
			sum += id
		}
	}
	return sum, nil
}

func SumGamePowers() (int, error) {
	id := "02-2"
	games, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var sum int
	for _, line := range games {
		sum += gamePower(&line)
	}
	return sum, nil
}

func validateGame(game *string) bool {
	start := nextIndexOfRune(game, ':')
	remaining := (*game)[start:]

	for len(remaining) > 0 {
		remaining = remaining[2:]

		nextSpace := nextIndexOfRune(&remaining, ' ')
		num := remaining[:nextSpace]
		count := parseInt(&num)
		char := remaining[nextSpace+1]

		switch char {
		case 'r':
			if count > redMaxCount {
				return false
			}
			remaining = remaining[nextSpace+4:]

		case 'g':
			if count > greenMaxCount {
				return false
			}
			remaining = remaining[nextSpace+6:]

		case 'b':
			if count > blueMaxCount {
				return false
			}
			remaining = remaining[nextSpace+5:]
		}
	}
	return true
}

func gamePower(game *string) int {
	start := nextIndexOfRune(game, ':')
	remaining := (*game)[start:]

	redMax := 0
	greenMax := 0
	blueMax := 0

	for len(remaining) > 0 {
		remaining = remaining[2:]

		nextSpace := nextIndexOfRune(&remaining, ' ')
		num := remaining[:nextSpace]
		count := parseInt(&num)
		char := remaining[nextSpace+1]

		switch char {
		case 'r':
			if count > redMax {
				redMax = count
			}
			remaining = remaining[nextSpace+4:]

		case 'g':
			if count > greenMax {
				greenMax = count
			}
			remaining = remaining[nextSpace+6:]

		case 'b':
			if count > blueMax {
				blueMax = count
			}
			remaining = remaining[nextSpace+5:]
		}
	}
	return redMax * greenMax * blueMax
}

func parseGameId(line *string) int {
	idx := nextIndexOfRune(line, ':')
	num := (*line)[5:idx]
	return parseInt(&num)
}

func parseInt(input *string) int {
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

func nextIndexOfRune(input *string, target rune) int {
	for i := range *input {
		if rune((*input)[i]) == target {
			return i
		}
	}
	return -1
}
