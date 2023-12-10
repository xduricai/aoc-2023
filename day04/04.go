package day04

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "04"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := getTotalPoints(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := getTotalCards(&lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)
	return nil
}

func getTotalPoints(cards *[]string) int {
	pointValues := []int{0, 1}
	var points int

	for _, card := range *cards {
		correct := countWinningNumbers(&card)

		for len(pointValues) <= correct {
			highest := pointValues[len(pointValues)-1]
			pointValues = append(pointValues, highest*2)
		}

		points += pointValues[correct]
	}

	return points
}

func getTotalCards(cards *[]string) int {
	var count int
	cardQuantities := make([]int, len(*cards))
	for i := range cardQuantities {
		cardQuantities[i] = 1
	}

	for idx, card := range *cards {
		correct := countWinningNumbers(&card)

		for i := 1; i <= correct; i++ {
			cardQuantities[idx+i] += cardQuantities[idx]
		}
		count += cardQuantities[idx]
	}

	return count
}

func countWinningNumbers(input *string) int {
	var correct int
	numMap := map[string]bool{}

	start := util.IndexOfRune(input, ':')
	separator := util.IndexOfRune(input, '|')

	winning := (*input)[start:separator]
	ours := (*input)[separator+1:]

	winningNums := util.ParseNumbersFromString(&winning)
	ourNums := util.ParseNumbersFromString(&ours)

	for _, num := range winningNums {
		numMap[num] = true
	}
	for _, num := range ourNums {
		if numMap[num] == true {
			correct++
		}
	}

	return correct
}
