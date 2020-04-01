package main

import (
	"bufio"
	"os"
)

func readLines(file string) ([]string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var lines []string
	s := bufio.NewScanner(fd)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err = s.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
