package day06

import (
	"math"

	"github.com/xduricai/aoc-2023/util"
)

func CountMultipleRaces() (int, error) {
	id := "06"
	input, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var total int

	times := util.ParseIntsFromString(&input[0])
	distances := util.ParseIntsFromString(&input[1])

	for idx := range times {
		a := 1.0
		b := float64(0 - times[idx])
		c := float64(distances[idx])

		x1, x2 := solveQuadraticEquation(a, b, c)
		res := int(math.Floor(x2)) - int(math.Floor(x1))

		if total > 0 {
			total *= res
		} else {
			total = res
		}
	}

	return total, nil
}

func CountSingleRace() (int, error) {
	id := "06"
	input, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var totalTime string
	var totalDistance string

	times := util.ParseNumbersFromString(&input[0])
	distances := util.ParseNumbersFromString(&input[1])

	for idx := range times {
		totalTime += times[idx]
		totalDistance += distances[idx]
	}

	a := 1.0
	b := float64(0 - util.ParseIntFromString(&totalTime))
	c := float64(util.ParseIntFromString(&totalDistance))

	x1, x2 := solveQuadraticEquation(a, b, c)
	res := int(math.Floor(x2)) - int(math.Floor(x1))
	return res, nil
}

// expects the discriminant to be non-negative
func solveQuadraticEquation(a float64, b float64, c float64) (float64, float64) {
	d := math.Pow(float64(b), 2) - (4 * a * c)

	x1 := (-b - math.Sqrt(d)) / (2 * a)
	x2 := (-b + math.Sqrt(d)) / (2 * a)

	return x1, x2
}
