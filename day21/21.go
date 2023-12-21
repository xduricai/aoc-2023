package day21

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "21"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := findTiles(&lines, 64)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := 0
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

var directions = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type State struct {
	x    int
	y    int
	dist int
}

func newState(x, y, dist int) *State {
	return &State{
		x:    x,
		y:    y,
		dist: dist,
	}
}

func findTiles(lines *[]string, steps int) int {
	var x int
	var y int
	height := len(*lines)
	width := len((*lines)[0])
	seen := map[State]bool{}

	for row := range *lines {
		for col := range (*lines)[row] {
			if (*lines)[row][col] == 'S' {
				x = col
				y = row
				break
			}
		}
	}

	inBounds := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height && (*lines)[y][x] != '#'
	}

	var walk func(int, int, int) int
	walk = func(x, y, dist int) int {
		if !inBounds(x, y) {
			return 0
		}

		state := *newState(x, y, dist)
		if seen[state] {
			return 0
		}
		seen[state] = true

		if dist >= steps {
			return 1
		}

		var total int
		for idx := range directions {
			total += walk(x+directions[idx][0], y+directions[idx][1], dist+1)
		}
		return total
	}
	return walk(x, y, 0)
}
