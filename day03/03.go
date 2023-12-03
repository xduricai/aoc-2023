package day03

import "github.com/xduricai/aoc-2023/util"

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

func SumAllParts() (int, error) {
	id := "03-1"
	lines, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var sum int
	height := len(lines)
	width := len(lines[0])

	for lineIdx, line := range lines {
		numStart := -1
		numEnd := -1
		validPart := false

		for idx, char := range line {
			if util.IsDigit(char) {
				if numStart < 0 {
					numStart = idx
				}
				numEnd = idx + 1

				neighbors := getNeighbors(lineIdx, idx, height, width)
				for _, p := range neighbors {
					c := rune(lines[p.row][p.col])
					if !util.IsDigit(c) && c != '.' {
						validPart = true
					}
				}
			} else {
				if !validPart {
					numStart = -1
					numEnd = -1
					continue
				}
				num := line[numStart:numEnd]
				sum += util.ParseIntFromString(&num)

				numStart = -1
				numEnd = -1
				validPart = false
			}
		}

		if validPart {
			num := line[numStart:numEnd]
			sum += util.ParseIntFromString(&num)
		}
	}
	return sum, nil
}

func SumGearRatios() (int, error) {
	id := "03-2"
	lines, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}

	var sum int
	height := len(lines)
	width := len(lines[0])

	for lineIdx, line := range lines {
		for idx := range line {
			if rune(line[idx]) != '*' {
				continue
			}
			numbers := []int{}

			if lineIdx > 0 {
				top := lines[lineIdx-1]

				if util.IsBDigit(top[idx]) {
					numbers = append(numbers, getNumber(&top, idx))
				} else {
					if idx-1 >= 0 && util.IsBDigit(top[idx-1]) {
						numbers = append(numbers, getNumber(&top, idx-1))
					}
					if idx+1 < width && util.IsBDigit(top[idx+1]) {
						numbers = append(numbers, getNumber(&top, idx+1))
					}
				}
			}

			if lineIdx+1 < height {
				bot := lines[lineIdx+1]

				if util.IsBDigit(bot[idx]) {
					numbers = append(numbers, getNumber(&bot, idx))
				} else {
					if idx-1 >= 0 && util.IsBDigit(bot[idx-1]) {
						numbers = append(numbers, getNumber(&bot, idx-1))
					}
					if idx+1 < width && util.IsBDigit(bot[idx+1]) {
						numbers = append(numbers, getNumber(&bot, idx+1))
					}
				}
			}

			if idx > 0 && util.IsBDigit(line[idx-1]) {
				numbers = append(numbers, getNumber(&line, idx-1))
			}
			if idx+1 < width && util.IsBDigit(line[idx+1]) {
				numbers = append(numbers, getNumber(&line, idx+1))
			}

			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}
	return sum, nil
}

func getNeighbors(row int, col int, height int, width int) []point {
	pts := []point{}

	isInBounds := func(row int, col int) bool {
		return row >= 0 && row < height && col >= 0 && col < width
	}

	if isInBounds(row-1, col-1) {
		pts = append(pts, *newPoint(row-1, col-1))
	}
	if isInBounds(row-1, col) {
		pts = append(pts, *newPoint(row-1, col))
	}
	if isInBounds(row-1, col+1) {
		pts = append(pts, *newPoint(row-1, col+1))
	}
	if isInBounds(row, col-1) {
		pts = append(pts, *newPoint(row, col-1))
	}
	if isInBounds(row, col+1) {
		pts = append(pts, *newPoint(row, col+1))
	}
	if isInBounds(row+1, col-1) {
		pts = append(pts, *newPoint(row+1, col-1))
	}
	if isInBounds(row+1, col) {
		pts = append(pts, *newPoint(row+1, col))
	}
	if isInBounds(row+1, col+1) {
		pts = append(pts, *newPoint(row+1, col+1))
	}

	return pts
}

func getNumber(line *string, start int) int {
	numStart := start
	numEnd := start + 1
	changed := true

	for changed {
		changed = false

		if numStart > 0 && util.IsBDigit((*line)[numStart-1]) {
			numStart--
			changed = true
		}
		if numEnd < len(*line) && util.IsBDigit((*line)[numEnd]) {
			numEnd++
			changed = true
		}
	}

	num := (*line)[numStart:numEnd]
	return util.ParseIntFromString(&num)
}
