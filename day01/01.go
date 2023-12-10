package day01

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "01"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := getNumericCoordinates(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := getMixedCoordinates(&lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)
	return nil
}

func getNumericCoordinates(coords *[]string) int {
	var tens int = 0
	var ones int = 0
	var val int = 0

	for _, line := range *coords {
		for i := range line {
			val = util.ParseIntFromRune(line[i])
			if val > 0 {
				tens += val
				break
			}
		}

		length := len(line)
		for i := range line {
			val = util.ParseIntFromRune(line[length-1-i])
			if val > 0 {
				ones += val
				break
			}
		}
	}

	return tens*10 + ones
}

func getMixedCoordinates(coords *[]string) int {
	var tens int = 0
	var ones int = 0
	var val int = 0

	for _, line := range *coords {
		for i := range line {
			val = parseNumberString(line[i:], false)
			if val > 0 {
				tens += val
				break
			}
		}

		length := len(line)
		for i := range line {
			val = parseNumberString(line[:length-i], true)
			if val > 0 {
				ones += val
				break
			}
		}
	}

	return tens*10 + ones
}

func parseNumberString(str string, reverse bool) int {
	length := len(str)
	var val int
	if reverse {
		val = util.ParseIntFromRune(str[length-1])
	} else {
		val = util.ParseIntFromRune(str[0])
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
