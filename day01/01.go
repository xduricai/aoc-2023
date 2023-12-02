package day01

import "github.com/xduricai/aoc-2023/util"

const RUNE_TO_DIGIT_OFFSET = 48

func GetNumericCoordinates() (int, error) {
	id := "01-1"
	coords, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var tens int = 0
	var ones int = 0
	var val int = 0

	for _, line := range coords {
		for i := range line {
			val = getCharacterValue(line[i])
			if val > 0 {
				tens += val
				break
			}
		}

		length := len(line)
		for i := range line {
			val = getCharacterValue(line[length-1-i])
			if val > 0 {
				ones += val
				break
			}
		}
	}

	return tens*10 + ones, nil
}

func GetMixedCoordinates() (int, error) {
	id := "01-2"
	coords, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var tens int = 0
	var ones int = 0
	var val int = 0

	for _, line := range coords {
		for i := range line {
			val = getSubstringValue(line[i:], false)
			if val > 0 {
				tens += val
				break
			}
		}

		length := len(line)
		for i := range line {
			val = getSubstringValue(line[:length-i], true)
			if val > 0 {
				ones += val
				break
			}
		}
	}

	return tens*10 + ones, nil
}

func getCharacterValue(b byte) int {
	r := rune(b)
	if r < '1' || r > '9' {
		return 0
	}

	return int(r) - RUNE_TO_DIGIT_OFFSET
}

func getSubstringValue(str string, reverse bool) int {
	length := len(str)
	var val int
	if reverse {
		val = getCharacterValue(str[length-1])
	} else {
		val = getCharacterValue(str[0])
	}

	if val > 0 {
		return val
	}
	if length < 3 {
		return 0
	}
	var word string

	if length >= 3 {
		if reverse {
			word = str[length-3:]
		} else {
			word = str[:3]
		}

		if word == "one" {
			return 1
		}
		if word == "two" {
			return 2
		}
		if word == "six" {
			return 6
		}
	}

	if length >= 4 {
		if reverse {
			word = str[length-4:]
		} else {
			word = str[:4]
		}

		if word == "four" {
			return 4
		}
		if word == "five" {
			return 5
		}
		if word == "nine" {
			return 9
		}
	}

	if length >= 5 {
		if reverse {
			word = str[length-5:]
		} else {
			word = str[:5]
		}

		if word == "three" {
			return 3
		}
		if word == "seven" {
			return 7
		}
		if word == "eight" {
			return 8
		}
	}

	return 0
}
