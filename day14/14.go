package day14

import (
	"fmt"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "14"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	// start1 := time.Now()
	fmt.Println(solve(&lines))
	// time1 := time.Since(start1)

	// start2 := time.Now()
	// part2 := solve(&lines, true)
	// time2 := time.Since(start2)

	// util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func solve(lines *[]string) int {
	var sum int
	height := len(*lines)
	width := len((*lines)[0])

	var nextVal int
	var colSum int

	for col := 0; col < width; col++ {
		nextVal = height
		colSum = 0

		for idx := 0; idx < height; idx++ {
			if (*lines)[idx][col] == '.' {
				continue
			}

			if (*lines)[idx][col] == 'O' {
				colSum += nextVal
				nextVal--
				continue
			} else {
				nextVal = height - idx - 1
			}
		}
		sum += colSum
	}
	return sum
}
