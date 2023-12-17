package day05

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "05"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	seeds, maps := parseData(&lines)

	start1 := time.Now()
	part1 := findLowestBasic(&seeds, &maps)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := findLowestExpanded(&seeds, &maps)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)
	return nil
}

func findLowestLocation(seed int, maps *[7][][]int) int {
	var num = seed

	for idx := 0; idx < 7; idx++ {
		for _, v := range (*maps)[idx] {
			if num >= v[1] && num < v[3] {
				num = num - v[1] + v[0]
				break
			}
		}
	}
	return num
}

func findLowestBasic(seeds *[]int, maps *[7][][]int) int {
	var min int
	var res int

	for _, seed := range *seeds {
		res = findLowestLocation(seed, maps)
		if res < min || min == 0 {
			min = res
		}
	}
	return min
}

func findLowestExpanded(seeds *[]int, maps *[7][][]int) int {
	var min int
	var res int

	for idx := 0; idx < len(*seeds); idx += 2 {
		for add := 0; add < (*seeds)[idx+1]; add++ {
			res = findLowestLocation((*seeds)[idx]+add, maps)
			if res < min || min == 0 {
				min = res
			}
		}
	}
	return min
}

func parseData(lines *[]string) ([]int, [7][][]int) {
	seeds := util.ParseIntsFromString(&(*lines)[0])
	maps := [7][][]int{}
	idx := -1

	for lIdx := 1; lIdx < len(*lines); lIdx++ {
		if len((*lines)[lIdx]) == 0 {
			lIdx++
			idx++
			continue
		}

		nums := util.ParseIntsFromString(&(*lines)[lIdx])
		nums = append(nums, nums[1]+nums[2])
		maps[idx] = append(maps[idx], nums)
	}
	return seeds, maps
}
