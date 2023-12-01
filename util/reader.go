package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(id string) ([]string, error) {
	path := fmt.Sprintf("../.inputs/%s.txt", id)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
