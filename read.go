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

func readLinesM(file string) (map[string]bool, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	m := make(map[string]bool)
	s := bufio.NewScanner(fd)
	for s.Scan() {
		m[s.Text()] = true
	}
	if err = s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}
