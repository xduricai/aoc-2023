package day17

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "17"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	height := len(lines)
	width := len(lines[0])

	input := make([][]int, height)
	for idx := range input {
		input[idx] = make([]int, width)
	}

	for row := range lines {
		for col := range lines[row] {
			input[row][col] = util.ParseIntFromRune(lines[row][col])
		}
	}

	start1 := time.Now()
	part1 := findPath(&input)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := findPath2(&input)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

var dirs = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func findPath(grid *[][]int) int {
	height := len(*grid)
	width := len((*grid)[0])

	inBounds := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height
	}

	seen := map[Move]bool{}
	heap := NewMinHeap()

	state := newState(0, newMove(0, 0, 1, 0))
	var move Move
	var x int
	var y int
	var dir int

	for ; state != nil; state = heap.Delete() {
		move = *state.move

		if move.x == width-1 && move.y == height-1 {
			return state.heat
		}
		if seen[move] {
			continue
		}
		seen[move] = true

		dir = move.dir
		x = move.x + dirs[dir][0]
		y = move.y + dirs[dir][1]

		if move.rep < 3 && inBounds(x, y) {
			heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, move.rep+1)))
		}

		dir = (dir + 1) % 4
		x = move.x + dirs[dir][0]
		y = move.y + dirs[dir][1]

		if inBounds(x, y) {
			heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, 1)))
		}

		dir = (dir + 2) % 4
		x = move.x + dirs[dir][0]
		y = move.y + dirs[dir][1]

		if inBounds(x, y) {
			heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, 1)))
		}
	}
	return 0
}

func findPath2(grid *[][]int) int {
	height := len(*grid)
	width := len((*grid)[0])

	inBounds := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height
	}

	seen := map[Move]bool{}
	heap := NewMinHeap()

	state := newState(0, newMove(0, 0, 1, 0))
	var move Move
	var x int
	var y int
	var dir int

	for ; state != nil; state = heap.Delete() {
		move = *state.move

		if move.x == width-1 && move.y == height-1 {
			return state.heat
		}
		if seen[move] {
			continue
		}
		seen[move] = true

		dir = move.dir
		x = move.x + dirs[dir][0]
		y = move.y + dirs[dir][1]

		if move.rep < 10 && inBounds(x, y) {
			heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, move.rep+1)))
		}

		if move.rep >= 4 {
			dir = (dir + 1) % 4
			x = move.x + dirs[dir][0]
			y = move.y + dirs[dir][1]

			if inBounds(x, y) {
				heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, 1)))
			}

			dir = (dir + 2) % 4
			x = move.x + dirs[dir][0]
			y = move.y + dirs[dir][1]

			if inBounds(x, y) {
				heap.Insert(newState(state.heat+(*grid)[y][x], newMove(x, y, dir, 1)))
			}
		}
	}
	return 0
}
