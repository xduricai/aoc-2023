package main

import (
	"fmt"
	"os"

	"github.com/xduricai/aoc-2023/day01"
	"github.com/xduricai/aoc-2023/day02"
	"github.com/xduricai/aoc-2023/day03"
	"github.com/xduricai/aoc-2023/day04"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify which day to run.")
		return
	}

	day := os.Args[1]
	var res1 any
	var res2 any
	var err1 error
	var err2 error

	switch day {
	case "1":
		res1, err1 = day01.GetNumericCoordinates()
		res2, err2 = day01.GetMixedCoordinates()
	case "2":
		res1, err1 = day02.SumValidGames()
		res2, err2 = day02.SumGamePowers()
	case "3":
		res1, err1 = day03.SumAllParts()
		res2, err2 = day03.SumGearRatios()
	case "4":
		res1, err1 = day04.GetTotalPoints()
		res2, err2 = day04.GetTotalCards()
	default:
		fmt.Printf("Day %s not recognized.", day)
		return
	}

	if err1 != nil {
		fmt.Println("An error occurred while running part 1.")
	}
	if err2 != nil {
		fmt.Println("An error occurred while running part 2.")
	}
	if err1 != nil || err2 != nil {
		return
	}

	fmt.Printf("DAY %s:\n", day)
	fmt.Printf("    Part 1: %v\n", res1)
	fmt.Printf("    Part 2: %v\n", res2)
}
