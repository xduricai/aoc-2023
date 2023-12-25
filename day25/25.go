package day25

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "25"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := splitGraph(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := 0
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func splitGraph(lines *[]string) int {
	graph := parseInput(lines)
	counts := map[string]int{}

	for start := range graph {
		seen := map[string]bool{}
		queue := util.NewQueue[string]()
		seen[start] = true
		queue.Enqueue(&start)

		for queue.Length() > 0 {
			src := *queue.Dequeue()

			for _, node := range graph[src] {
				if seen[node] {
					continue
				}
				seen[node] = true

				counts[edgeName(&src, &node)]++
				queue.Enqueue(&node)
			}
		}
	}

	max := getMax(&counts)
	res := 1

	for _, start := range [2]string{max[0][0:3], max[0][3:]} {
		seen := map[string]bool{}
		queue := util.NewQueue[string]()
		count := 1
		seen[start] = true
		queue.Enqueue(&start)

		for queue.Length() > 0 {
			src := *queue.Dequeue()

			for _, node := range graph[src] {
				edge := edgeName(&src, &node)
				if edge == max[0] || edge == max[1] || edge == max[2] {
					continue
				}

				if seen[node] {
					continue
				}
				seen[node] = true

				queue.Enqueue(&node)
				count++
			}
		}
		res *= count
	}

	return res
}

func getMax(counts *map[string]int) [3]string {
	maxV := [3]int{}
	maxK := [3]string{}

	for edge := range *counts {
		if (*counts)[edge] > maxV[0] {
			maxV[2] = maxV[1]
			maxK[2] = maxK[1]

			maxV[1] = maxV[0]
			maxK[1] = maxK[0]

			maxV[0] = (*counts)[edge]
			maxK[0] = edge
		} else if (*counts)[edge] > maxV[1] {
			maxV[2] = maxV[1]
			maxK[2] = maxK[1]

			maxV[1] = (*counts)[edge]
			maxK[1] = edge
		} else if (*counts)[edge] > maxV[2] {
			maxV[2] = (*counts)[edge]
			maxK[2] = edge
		}
	}
	return maxK
}

func edgeName(a, b *string) string {
	if *a < *b {
		return *a + *b
	}
	return *b + *a
}

func parseInput(lines *[]string) map[string][]string {
	graph := map[string][]string{}

	for _, line := range *lines {
		src := line[:3]

		for idx, char := range line {
			if char == ' ' {
				edge := line[idx+1 : idx+4]
				graph[src] = append(graph[src], edge)
				graph[edge] = append(graph[edge], src)
			}
		}
	}
	return graph
}
