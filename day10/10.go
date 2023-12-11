package day10

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "10"
	maze, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1, path, vertrices := findLoop(&maze)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := countInnerTiles(&vertrices, &path)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

// 0 = UP, 1 = RIGHT, 2 = DOWN, 3 = LEFT
var directions = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func findLoop(maze *[]string) (int, [][2]int, [][2]int) {
	var startX, startY int
	height := len(*maze)
	width := len((*maze)[0])

	for row, line := range *maze {
		col := util.IndexOfRune(&line, 'S')
		if col != -1 {
			startX = col
			startY = row
			break
		}
	}

	path := [][2]int{}
	vertrices := [][2]int{}
	for i := range directions {
		path = [][2]int{}
		vertrices = [][2]int{}
		x := startX
		y := startY

		if !inBounds(x+directions[i][0], y+directions[i][1], width, height) {
			continue
		}

		dir := i
		current := 'X'

		for dir != -1 && current != 'S' {
			dir = nextDirection(rune(current), dir)
			x += directions[dir][0]
			y += directions[dir][1]
			current = rune((*maze)[y][x])

			if current == 'F' || current == '7' || current == 'L' || current == 'J' {
				vertrices = append(vertrices, [2]int{x, y})
			}
			path = append(path, [2]int{x, y})
		}

		if current == 'S' {
			if len(vertrices)%2 == 1 {
				vertrices = append(vertrices, [2]int{x, y})
			}
			return len(path) / 2, path, vertrices
		}
	}
	return 0, nil, nil
}

func countInnerTiles(vertrices *[][2]int, path *[][2]int) int {
	var area int

	previous := (*vertrices)[len(*vertrices)-1]
	current := previous

	for _, point := range *vertrices {
		previous = current
		current = point
		area += ((previous[1] * current[0]) - (previous[0] * current[1]))
	}
	area /= 2

	return area - (len(*path) / 2) + 1
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}

func nextDirection(c rune, dir int) int {

	switch {
	case c == '|' && dir == 0:
		return 0
	case c == '|' && dir == 2:
		return 2
	case c == '-' && dir == 1:
		return 1
	case c == '-' && dir == 3:
		return 3
	case c == 'F' && dir == 0:
		return 1
	case c == 'F' && dir == 3:
		return 2
	case c == '7' && dir == 0:
		return 3
	case c == '7' && dir == 1:
		return 2
	case c == 'L' && dir == 2:
		return 1
	case c == 'L' && dir == 3:
		return 0
	case c == 'J' && dir == 1:
		return 0
	case c == 'J' && dir == 2:
		return 3
	case c == 'X':
		return dir
	default:
		return -1
	}
}
