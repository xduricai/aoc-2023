package day18

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "18"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := calculateArea(&lines, false)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := calculateArea(&lines, true)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

type Move struct {
	dist int
	x    int
	y    int
	dir  byte
}

type Point struct {
	x int64
	y int64
}

func newPoint(x, y int) *Point {
	return &Point{
		x: int64(x),
		y: int64(y),
	}
}

func calculateArea(lines *[]string, colors bool) int64 {
	var moves []Move
	if colors {
		moves = parseColors(lines)
	} else {
		moves = parsDirections(lines)
	}

	vertrices := []Point{}
	dir := byte('X')
	var x int
	var y int
	var circ int64

	for _, move := range moves {
		x += move.x
		y += move.y
		circ += int64(move.dist)

		if dir != move.dir {
			dir = move.dir
			vertrices = append(vertrices, *newPoint(x, y))
		}
	}

	area := circ
	prev := vertrices[len(vertrices)-1]
	curr := prev

	for _, point := range vertrices {
		prev = curr
		curr = point
		area += (curr.y*prev.x - (curr.x * prev.y))
	}
	area /= 2

	return area + 1
}

func parsDirections(lines *[]string) []Move {
	moves := make([]Move, len(*lines))
	var sep int
	var num int
	var str string

	for idx, line := range *lines {
		sep = util.IndexOfRune(&line, '(')
		str = line[:sep-1]
		num = util.ParseIntFromString(&str)

		moves[idx].dist = num
		moves[idx].dir = line[0]

		switch line[0] {
		case 'U':
			moves[idx].x = 0
			moves[idx].y = -num
		case 'R':
			moves[idx].x = num
			moves[idx].y = 0
		case 'D':
			moves[idx].x = 0
			moves[idx].y = num
		case 'L':
			moves[idx].x = -num
			moves[idx].y = 0
		}
	}
	return moves
}

func parseColors(lines *[]string) []Move {
	moves := make([]Move, len((*lines)))
	var sep int
	var num int
	var color string

	for idx, line := range *lines {
		sep = util.IndexOfRune(&line, '#')
		color = line[sep+1 : len(line)-2]
		num = util.ParseIntFromHexString(&color)

		moves[idx].dist = num
		switch line[len(line)-2] {
		case '3':
			moves[idx].dir = 'U'
			moves[idx].x = 0
			moves[idx].y = -num
		case '0':
			moves[idx].dir = 'R'
			moves[idx].x = num
			moves[idx].y = 0
		case '1':
			moves[idx].dir = 'D'
			moves[idx].x = 0
			moves[idx].y = num
		case '2':
			moves[idx].dir = 'L'
			moves[idx].x = -num
			moves[idx].y = 0
		}
	}
	return moves
}
