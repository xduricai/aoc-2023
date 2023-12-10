package day08

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

type Node struct {
	left  string
	right string
}

func newNode(left *string, right *string) *Node {
	return &Node{
		left:  *left,
		right: *right,
	}
}

func Run() error {
	id := "08"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := findSinglePath(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := FindMultiplePaths(&lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func findSinglePath(input *[]string) int {
	path := (*input)[0]
	lines := (*input)[2:]
	nodeMap := map[string]Node{}

	for _, line := range lines {
		tag := line[:util.IndexOfRune(&line, ' ')]
		left := line[util.IndexOfRune(&line, '(')+1 : util.IndexOfRune(&line, ',')]
		right := line[util.IndexOfRune(&line, ',')+2 : util.IndexOfRune(&line, ')')]
		nodeMap[tag] = *newNode(&left, &right)
	}

	start := "AAA"
	isTarget := func(node *string) bool {
		return *node == "ZZZ"
	}

	return findPath(&start, &path, &nodeMap, isTarget)
}

func FindMultiplePaths(input *[]string) int {
	path := (*input)[0]
	lines := (*input)[2:]
	nodeMap := map[string]Node{}
	starts := []string{}
	results := []int{}

	for _, line := range lines {
		tag := line[:util.IndexOfRune(&line, ' ')]
		left := line[util.IndexOfRune(&line, '(')+1 : util.IndexOfRune(&line, ',')]
		right := line[util.IndexOfRune(&line, ',')+2 : util.IndexOfRune(&line, ')')]
		nodeMap[tag] = *newNode(&left, &right)

		if tag[2] == 'A' {
			starts = append(starts, tag)
		}
	}

	isTarget := func(node *string) bool {
		return (*node)[2] == 'Z'
	}

	for _, start := range starts {
		results = append(results, findPath(&start, &path, &nodeMap, isTarget))
	}

	count := len(results)
	if count == 1 {
		return results[0]
	}
	return util.LCM(results[0], results[1], results[2:]...)
}

func findPath(start *string, path *string, nodeMap *map[string]Node, isTarget func(*string) bool) int {
	var steps int
	idx := 0
	current := *start
	pathLen := len(*path)

	for !isTarget(&current) {
		if (*path)[idx] == 'L' {
			current = (*nodeMap)[current].left
		} else {
			current = (*nodeMap)[current].right
		}

		idx = (idx + 1) % pathLen
		steps++
	}

	return steps
}
