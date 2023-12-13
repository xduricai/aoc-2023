package util

import (
	"fmt"
	"time"
)

func PrintResults(day string, res1, res2 any, time1, time2 time.Duration) {
	fmt.Printf("DAY %s:\n", day)
	fmt.Printf("    Part 1: %v in %v\n", res1, time1)
	fmt.Printf("    Part 2: %v in %v\n", res2, time2)
}
