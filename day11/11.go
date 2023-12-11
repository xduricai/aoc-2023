package day11

import (
	"math"
	"time"

	"github.com/xduricai/aoc-2023/util"
)

type point struct {
	row int
	col int
}

func newPoint(row int, col int) *point {
	return &point{
		row: row,
		col: col,
	}
}

func Run() error {
	id := "11"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := getDistances(1, &lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := getDistances(999999, &lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func getDistances(increment int, lines *[]string) int {
	var sum int
	height := len(*lines)
	width := len((*lines)[0])

	galaxyRows := make([]bool, height)
	galaxyCols := make([]bool, width)
	galaxies := []point{}

	for row := range *lines {
		for col := range (*lines)[row] {
			if rune((*lines)[row][col]) == '#' {
				galaxies = append(galaxies, *newPoint(row, col))
				galaxyRows[row] = true
				galaxyCols[col] = true
			}
		}
	}

	rowExpansions := make([]int, height)
	colExpansions := make([]int, width)

	if !galaxyRows[0] {
		rowExpansions[0] = increment
	}
	for idx := 1; idx < height; idx++ {
		rowExpansions[idx] = rowExpansions[idx-1]
		if !galaxyRows[idx] {
			rowExpansions[idx] += increment
		}
	}

	if !galaxyCols[0] {
		colExpansions[0] = increment
	}
	for idx := 1; idx < height; idx++ {
		if !galaxyCols[idx] {
			colExpansions[idx] = increment
		}
		colExpansions[idx] += colExpansions[idx-1]
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			sum += distance(
				galaxies[i].row+rowExpansions[galaxies[i].row],
				galaxies[i].col+colExpansions[galaxies[i].col],
				galaxies[j].row+rowExpansions[galaxies[j].row],
				galaxies[j].col+colExpansions[galaxies[j].col],
			)
		}
	}

	return sum
}

func distance(row1, col1, row2, col2 int) int {
	return int(math.Abs(float64(row1)-float64(row2)) + math.Abs(float64(col1)-float64(col2)))
}
