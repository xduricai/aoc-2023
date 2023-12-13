package day12

import (
	"fmt"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "12"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	part1 := sumArrangements(&lines)
	fmt.Println(part1)

	return nil
}

func sumArrangements(lines *[]string) int {
	var sum int
	return sum
}

func parseLine(line *string) (string, []int) {
	return (*line)[:util.IndexOfRune(line, ' ')], util.ParseIntsFromString(&(*line))
}
