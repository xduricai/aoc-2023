package day02

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

const redMaxCount = 12
const greenMaxCount = 13
const blueMaxCount = 14

func Run() error {
	id := "02"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start := time.Now()
	part1, part2 := sumGames(&lines)
	time1 := time.Since(start)

	util.PrintResults(id, part1, part2, time1, time1)
	return nil
}

func sumGames(games *[]string) (int, int) {
	var gamesSum int
	var powersSum int
	for _, line := range *games {
		id := parseGameId(&line)
		if validateGame(&line) {
			gamesSum += id
		}
		powersSum += gamePower(&line)
	}
	return gamesSum, powersSum
}

func validateGame(game *string) bool {
	start := util.IndexOfRune(game, ':')
	remaining := (*game)[start:]

	for len(remaining) > 0 {
		remaining = remaining[2:]

		nextSpace := util.IndexOfRune(&remaining, ' ')
		num := remaining[:nextSpace]
		count := util.ParseIntFromString(&num)
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
	start := util.IndexOfRune(game, ':')
	remaining := (*game)[start:]

	redMax := 0
	greenMax := 0
	blueMax := 0

	for len(remaining) > 0 {
		remaining = remaining[2:]

		nextSpace := util.IndexOfRune(&remaining, ' ')
		num := remaining[:nextSpace]
		count := util.ParseIntFromString(&num)
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
	idx := util.IndexOfRune(line, ':')
	num := (*line)[5:idx]
	return util.ParseIntFromString(&num)
}
