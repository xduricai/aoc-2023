package day16

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "16"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := countActive(&lines, 0, 0, 1)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := findOptimal(&lines)
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

func countActive(lines *[]string, startX int, startY int, startDir int) int {
	height := len(*lines)
	width := len((*lines)[0])

	visited := make([][][4]bool, height)
	for idx := range visited {
		visited[idx] = make([][4]bool, width)
	}

	inBounds := func(x int, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height
	}

	var walk func(int, int, int) int
	walk = func(x, y, dir int) int {
		var total int

		for inBounds(x, y) {
			if visited[y][x][dir] {
				break
			}
			if !visited[y][x][0] && !visited[y][x][1] && !visited[y][x][2] && !visited[y][x][3] {
				total++
			}
			visited[y][x][dir] = true

			dir = getDirection((*lines)[y][x], dir)

			if dir == 4 {
				total += walk(x+1, y, 1)
				total += walk(x-1, y, 3)
				break
			}
			if dir == 5 {
				total += walk(x, y-1, 0)
				total += walk(x, y+1, 2)
				break
			}

			x += directions[dir][0]
			y += directions[dir][1]
		}
		return total
	}
	return walk(startX, startY, startDir)
}

func findOptimal(lines *[]string) int {
	var max int
	var current int
	height := len(*lines)
	width := len((*lines)[0])

	for idx := 0; idx < width; idx++ {
		current = countActive(lines, idx, 0, 2)
		if current > max {
			max = current
		}
		current = countActive(lines, idx, height-1, 0)
		if current > max {
			max = current
		}
	}

	for idx := 0; idx < height; idx++ {
		current = countActive(lines, 0, idx, 1)
		if current > max {
			max = current
		}
		current = countActive(lines, width-1, idx, 3)
		if current > max {
			max = current
		}
	}

	return max
}

func getDirection(char byte, dir int) int {
	switch {
	case char == '\\' && dir == 0:
		return 3
	case char == '\\' && dir == 1:
		return 2
	case char == '\\' && dir == 2:
		return 1
	case char == '\\' && dir == 3:
		return 0
	case char == '/' && dir == 0:
		return 1
	case char == '/' && dir == 1:
		return 0
	case char == '/' && dir == 2:
		return 3
	case char == '/' && dir == 3:
		return 2
	case char == '-' && (dir == 0 || dir == 2):
		return 4
	case char == '|' && (dir == 1 || dir == 3):
		return 5
	default:
		return dir
	}
}
