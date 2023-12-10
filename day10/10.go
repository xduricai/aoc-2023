package day10

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

// 0 = UP, 1 = RIGHT, 2 = DOWN, 3 = LEFT
var directions = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var movesTo = [4][4]rune{
	{'F', '7', '|', 'S'},
	{'7', 'J', '-', 'S'},
	{'L', 'J', '|', 'S'},
	{'F', 'L', '-', 'S'},
}

var movesFrom = [7][]int{
	{0, 2},
	{1, 3},
	{1, 2},
	{2, 3},
	{0, 1},
	{0, 3},
	{0, 1, 2, 3},
}

func Run() error {
	id := "10"
	maze, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	part1, part2, time1, time2 := solve(&maze)
	util.PrintResults(id, part1, part2, time1, time2)
	return nil
}

func solve(maze *[]string) (int, int, time.Duration, time.Duration) {
	var startX, startY int
	height := len(*maze)
	width := len((*maze)[0])
	visited := make([][]bool, height)

	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	for row, line := range *maze {
		col := util.IndexOfRune(&line, 'S')
		if col != -1 {
			startX = col
			startY = row
			break
		}
	}
	path := [][2]int{}

	start1 := time.Now()
	dfs(startX, startY, 0, &path, maze, &visited)
	farthest := len(path) / 2
	time1 := time.Since(start1)

	start2 := time.Now()
	inner := countInnerTiles(maze, &path)
	time2 := time.Since(start2)

	return farthest, inner, time1, time2
}

func countInnerTiles(maze *[]string, path *[][2]int) int {
	var area int
	vertrices := [][2]int{}

	var point [2]int
	for idx := 1; idx < len(*path); idx++ {
		point = (*path)[idx]

		if idx == 0 || (*maze)[point[1]][point[0]] == '|' || (*maze)[point[1]][point[0]] == '-' {
			continue
		}
		vertrices = append(vertrices, point)
	}

	if len(vertrices)%2 == 1 {
		vertrices = append(vertrices, (*path)[0])
	}

	previous := vertrices[len(vertrices)-1]
	current := previous

	for _, point := range vertrices {
		previous = current
		current = point
		area += ((previous[1] * current[0]) - (previous[0] * current[1]))
	}
	area /= 2

	return area - (len(*path) / 2) + 1
}

func getMoveIndex(char rune) int {
	switch char {
	case '|':
		return 0
	case '-':
		return 1
	case 'F':
		return 2
	case '7':
		return 3
	case 'L':
		return 4
	case 'J':
		return 5
	default:
		return 6
	}
}

func dfs(x, y, direction int, path *[][2]int, maze *[]string, visited *[][]bool) bool {
	// out of bounds check
	if x < 0 || y < 0 || x >= len((*maze)[0]) || y >= len(*maze) {
		return false
	}

	current := rune((*maze)[y][x])
	// no path
	if current == '.' {
		return false
	}
	// loop found
	if current == 'S' && len(*path) > 1 {
		return true
	}
	if (*visited)[y][x] {
		return false
	}

	valid := false
	for _, move := range movesTo[direction] {
		if move == current {
			valid = true
			break
		}
	}
	if !valid {
		return false
	}

	(*visited)[y][x] = true
	*path = append(*path, [2]int{x, y})

	for _, idx := range movesFrom[getMoveIndex(current)] {
		dir := directions[idx]

		result := dfs(x+dir[0], y+dir[1], idx, path, maze, visited)
		if result {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]

	return false
}
