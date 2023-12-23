package day23

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "23"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := findLongestPath(&lines, true)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := findLongestPath(&lines, false)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

var directions = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Point [2]int
type Graph map[Point][]Edge
type Edge struct {
	x    int
	y    int
	dist int
}

func newPoint(x, y int) *Point {
	return &Point{x, y}
}

func newGraph() *Graph {
	return &Graph{}
}

func newEdge(x, y, dist int) *Edge {
	return &Edge{
		x:    x,
		y:    y,
		dist: dist,
	}
}

func findLongestPath(maze *[]string, slopes bool) int {
	height := len(*maze)
	width := len((*maze)[0])

	var start, end Point
	for idx, char := range (*maze)[0] {
		if char == '.' {
			start = *newPoint(idx, 0)
		}
	}
	for idx, char := range (*maze)[height-1] {
		if char == '.' {
			end = *newPoint(idx, height-1)
		}
	}

	nodes := map[Point]bool{
		start: true,
		end:   true,
	}

	isValid := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height && (*maze)[y][x] != '#'
	}

	var neighbors int
	for y, line := range *maze {
		for x := range line {
			neighbors = 0
			if (*maze)[y][x] == '#' {
				continue
			}

			for _, dir := range directions {
				if isValid(x+dir[0], y+dir[1]) {
					neighbors++
				}
			}
			if neighbors > 2 {
				nodes[*newPoint(x, y)] = true
			}
		}
	}

	graph := *newGraph()

	for node := range nodes {
		seen := map[Point]bool{node: true}
		stack := util.NewStack[Edge]()
		stack.Push(*newEdge(node[0], node[1], 0))

		for edge := stack.Pop(); edge != nil; edge = stack.Pop() {
			point := *newPoint(edge.x, edge.y)

			if nodes[point] && edge.dist > 0 {
				graph[node] = append(graph[node], *edge)
				continue
			}

			for _, dir := range getDirs((*maze)[edge.y][edge.x], slopes) {
				x := edge.x + dir[0]
				y := edge.y + dir[1]
				point[0] = x
				point[1] = y

				if isValid(x, y) && !seen[point] {
					stack.Push(*newEdge(x, y, edge.dist+1))
					seen[*newPoint(x, y)] = true
				}
			}
		}
	}

	seen := map[Point]bool{}
	var walk func(Point) int

	walk = func(point Point) int {
		if point == end {
			return 0
		}
		maxDist := util.MinInt32

		seen[point] = true
		for _, node := range graph[point] {
			next := *newPoint(node.x, node.y)

			if seen[next] {
				continue
			}
			maxDist = max(maxDist, walk(next)+node.dist)
		}
		seen[point] = false

		return maxDist
	}

	return walk(start)
}

func getDirs(tile byte, withSlopes bool) [][2]int {
	if withSlopes && tile != '.' {
		switch tile {
		case '^':
			return directions[0:1]
		case '>':
			return directions[1:2]
		case 'v':
			return directions[2:3]
		case '<':
			return directions[3:4]
		}
	}
	return directions[:]
}
