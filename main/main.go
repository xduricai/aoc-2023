package main

import (
	"fmt"
	"os"

	"github.com/xduricai/aoc-2023/day01"
	"github.com/xduricai/aoc-2023/day02"
	"github.com/xduricai/aoc-2023/day03"
	"github.com/xduricai/aoc-2023/day04"
	"github.com/xduricai/aoc-2023/day05"
	"github.com/xduricai/aoc-2023/day06"
	"github.com/xduricai/aoc-2023/day07"
	"github.com/xduricai/aoc-2023/day08"
	"github.com/xduricai/aoc-2023/day09"
	"github.com/xduricai/aoc-2023/day10"
	"github.com/xduricai/aoc-2023/day11"
	"github.com/xduricai/aoc-2023/day12"
	"github.com/xduricai/aoc-2023/day13"
	"github.com/xduricai/aoc-2023/day14"
	"github.com/xduricai/aoc-2023/day15"
	"github.com/xduricai/aoc-2023/day16"
	"github.com/xduricai/aoc-2023/day17"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify which day to run.")
		return
	}

	day := os.Args[1]
	var err error

	switch day {
	case "1":
		err = day01.Run()
	case "2":
		err = day02.Run()
	case "3":
		err = day03.Run()
	case "4":
		err = day04.Run()
	case "5":
		err = day05.Run()
	case "6":
		err = day06.Run()
	case "7":
		err = day07.Run()
	case "8":
		err = day08.Run()
	case "9":
		err = day09.Run()
	case "10":
		err = day10.Run()
	case "11":
		err = day11.Run()
	case "12":
		err = day12.Run()
	case "13":
		err = day13.Run()
	case "14":
		err = day14.Run()
	case "15":
		err = day15.Run()
	case "16":
		err = day16.Run()
	case "17":
		err = day17.Run()
	// case "18":
	// 	err = day18.Run()
	// case "19":
	// 	err = day19.Run()
	// case "20":
	// 	err = day20.Run()
	// case "21":
	// 	err = day21.Run()
	// case "22":
	// 	err = day22.Run()
	// case "23":
	// 	err = day23.Run()
	// case "24":
	// 	err = day24.Run()
	// case "25":
	// 	err = day25.Run()
	default:
		fmt.Printf("Day %s not recognized.", day)
		return
	}

	if err != nil {
		fmt.Println("An error occurred while loading data.")
	}
}
