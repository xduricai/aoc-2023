package day10

import "github.com/xduricai/aoc-2023/util"

func Run() (int, int, error) {
	id := "10"
	_, err := util.ReadLines(id)

	if err != nil {
		return *new(int), *new(int), err
	}

	return 0, 0, nil
}
