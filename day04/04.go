package day04

import "github.com/xduricai/aoc-2023/util"

func GetTotalPoints() (int, error) {
	id := "04-1"
	cards, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	pointValues := []int{0, 1}
	var points int

	for _, card := range cards {
		correct := countWinningNumbers(&card)

		for len(pointValues) <= correct {
			highest := pointValues[len(pointValues)-1]
			pointValues = append(pointValues, highest*2)
		}

		points += pointValues[correct]
	}

	return points, nil
}

func GetTotalCards() (int, error) {
	id := "04-2"
	cards, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var count int
	cardQuantities := make([]int, len(cards))
	for i := range cardQuantities {
		cardQuantities[i] = 1
	}

	for idx, card := range cards {
		correct := countWinningNumbers(&card)

		for i := 1; i <= correct; i++ {
			cardQuantities[idx+i] += cardQuantities[idx]
		}
		count += cardQuantities[idx]
	}

	return count, nil
}

func countWinningNumbers(input *string) int {
	var correct int
	numMap := map[string]bool{}

	start := util.IndexOfRune(input, ':')
	separator := util.IndexOfRune(input, '|')

	winning := (*input)[start:separator]
	ours := (*input)[separator+1:]

	winningNums := parseNumbers(&winning)
	ourNums := parseNumbers(&ours)

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

func parseNumbers(input *string) []string {
	numbers := []string{}

	numStart := -1
	numEnd := -1

	for idx, char := range *input {
		if util.IsDigit(char) {
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
