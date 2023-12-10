package day06

import (
	"math"
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "06"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := countRaces(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := countSingleRace(&lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func countRaces(input *[]string) int {
	var total int

	times := util.ParseIntsFromString(&(*input)[0])
	distances := util.ParseIntsFromString(&(*input)[1])

	for idx := range times {
		a := 1.0
		b := float64(0 - times[idx])
		c := float64(distances[idx])

		x1, x2 := util.SolveQuadraticEquation(a, b, c)
		res := int(math.Floor(x2)) - int(math.Floor(x1))

		if total > 0 {
			total *= res
		} else {
			total = res
		}
	}

	return total
}

func countSingleRace(input *[]string) int {
	var totalTime string
	var totalDistance string

	times := util.ParseNumbersFromString(&(*input)[0])
	distances := util.ParseNumbersFromString(&(*input)[1])

	for idx := range times {
		totalTime += times[idx]
		totalDistance += distances[idx]
	}

	a := 1.0
	b := float64(0 - util.ParseIntFromString(&totalTime))
	c := float64(util.ParseIntFromString(&totalDistance))

	x1, x2 := util.SolveQuadraticEquation(a, b, c)
	res := int(math.Floor(x2)) - int(math.Floor(x1))
	return res
}
