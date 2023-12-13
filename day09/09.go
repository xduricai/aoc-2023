package day09

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "09"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := extrapolateValues(&lines, true)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := extrapolateValues(&lines, false)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func extrapolateValues(lines *[]string, newValues bool) int {
	var sum int

	for _, line := range *lines {
		values := [][]int{util.ParseIntsFromString(&line)}
		valueChanged := true

		for i := 1; valueChanged; i++ {
			valueChanged = false
			length := len(values[i-1]) - 1

			values = append(values, make([]int, length))
			values[i][0] = values[i-1][1] - values[i-1][0]

			for idx := 1; idx < length; idx++ {
				values[i][idx] = values[i-1][idx+1] - values[i-1][idx]

				if values[i][idx] != values[i][idx-1] {
					valueChanged = true
				}
			}
		}

		var final int
		if newValues {
			idx := len(values[0])
			for _, row := range values {
				idx--
				final += row[idx]
			}
		} else {
			negative := false

			for _, row := range values {
				if negative {
					final -= row[0]
				} else {
					final += row[0]
				}
				negative = !negative
			}
		}

		sum += final
	}

	return sum
}
