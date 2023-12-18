package day19

import (
	"fmt"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "19"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	// start1 := time.Now()
	part1 := solve(&lines)
	fmt.Println(part1)
	// time1 := time.Since(start1)

	// start2 := time.Now()
	// part2 := solve(&lines)
	// time2 := time.Since(start2)

	// util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func solve(lines *[]string) int {
	var sum int

	return sum
}
