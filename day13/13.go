package day13

import (
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "13"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := solve(&lines, false)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := solve(&lines, true)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil
}

func solve(lines *[]string, smudges bool) int {
	blocks := parseBlocks(lines)
	var sum int

	for _, block := range blocks {
		res := 100 * manacher(block, smudges)
		if res == 0 {
			res = manacher(util.Transpose(&block), smudges)
		}
		sum += res
	}
	return sum
}

func padInput(input *[]string) []string {
	output := make([]string, len(*input)*2+1)
	output[0] = ""

	idx := 1
	for _, char := range *input {
		output[idx] = char
		idx++
		output[idx] = ""
		idx++
	}
	return output
}

func manacher(input []string, smudges bool) int {
	lines := padInput(&input)
	n := len(lines)
	p := make([]int, n)

	var center int
	var right int
	var maxLen int
	var maxIdx int

	for idx := 1; idx < n-1; idx++ {
		if right > idx {
			mirror := 2*center - idx
			p[idx] = min(right-idx, p[mirror])
		}

		l := idx - p[idx] - 1
		r := idx + p[idx] + 1

		if !smudges {
			for l >= 0 && r < n && lines[l] == lines[r] {
				p[idx]++
				l = idx - p[idx] - 1
				r = idx + p[idx] + 1
			}
		} else {
			dirty := false
			initial := p[idx]

			for l >= 0 && r < n {
				equal, smudge := compareLines(&lines[l], &lines[r])
				if !equal || (dirty && smudge) {
					break
				}
				if smudge {
					dirty = true
				}

				p[idx]++
				l = idx - p[idx] - 1
				r = idx + p[idx] + 1
			}
			if !dirty {
				p[idx] = initial
			}
		}

		if idx+p[idx] > right {
			center = idx
			right = idx + p[idx]
		}

		if p[idx] > maxLen && p[idx]%2 == 0 && (idx-p[idx] == 0 || idx+p[idx] == n-1) {
			maxLen = p[idx]
			maxIdx = idx
		}
	}
	return maxIdx / 2
}

func parseBlocks(input *[]string) [][]string {
	output := [][]string{}
	current := []string{}

	for _, line := range *input {
		if len(line) > 0 {
			current = append(current, line)
			continue
		}
		output = append(output, current)
		current = []string{}
	}
	output = append(output, current)

	return output
}

func compareLines(left *string, right *string) (bool, bool) {
	var diffs int
	for idx := range *left {
		if (*left)[idx] != (*right)[idx] {
			diffs++
		}
		if diffs > 1 {
			return false, false
		}
	}
	return true, diffs == 1
}
