package day15

import (
	"strings"
	"time"

	"github.com/xduricai/aoc-2023/util"
)

func Run() error {
	id := "15"
	lines, err := util.ReadLines(id)

	if err != nil {
		return err
	}

	start1 := time.Now()
	part1 := sumHashes(&lines)
	time1 := time.Since(start1)

	start2 := time.Now()
	part2 := sumLensPowers(&lines)
	time2 := time.Since(start2)

	util.PrintResults(id, part1, part2, time1, time2)

	return nil

}

func sumHashes(lines *[]string) int {
	var sum int
	var current int

	for _, line := range *lines {
		for _, char := range line {
			if char == ',' {
				sum += current
				current = 0
				continue
			}

			current += int(char)
			current *= 17
			current %= 256
		}
	}
	sum += current

	return sum
}

func sumLensPowers(lines *[]string) int {
	var sum int
	boxes := make([]LinkedList, 256)

	var sb strings.Builder
	var label string
	var hash int
	var value int

	for _, line := range *lines {
		for _, char := range line {
			if char == ',' {
				label, hash, value = parseInstruction(sb.String())
				if value >= 0 {
					boxes[hash].Insert(&label, value)
				} else {
					boxes[hash].Remove(&label)
				}
				sb.Reset()
			} else {
				sb.WriteRune(char)
			}
		}
	}

	label, hash, value = parseInstruction(sb.String())
	if value >= 0 {
		boxes[hash].Insert(&label, value)
	} else {
		boxes[hash].Remove(&label)
	}

	var idx int
	var box LinkedList
	var current *Node

	for boxIdx := 1; boxIdx <= 256; boxIdx++ {
		box = boxes[boxIdx-1]

		if box.length == 0 {
			continue
		}

		current = box.head
		idx = 1

		for current != nil {
			sum += (boxIdx * idx * current.value)
			current = current.next
			idx++
		}
	}

	return sum
}

func parseInstruction(input string) (string, int, int) {
	var hash int

	for idx := range input {
		if input[idx] == '-' {
			return input[:idx], hash, -1
		}
		if input[idx] == '=' {
			num := input[idx+1:]
			return input[:idx], hash, util.ParseIntFromString(&num)
		}

		hash += int(input[idx])
		hash *= 17
		hash %= 256
	}
	return "", -1, -1
}
