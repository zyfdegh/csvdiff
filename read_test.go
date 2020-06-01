package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	lines, err := readLines("testdata/a.csv")
	assert.Nil(t, err)
	assert.NotNil(t, lines)
	assert.Len(t, lines, 3)
}

func TestReadLinesM(t *testing.T) {
	m, err := readLines("testdata/a.csv")
	assert.Nil(t, err)
	assert.NotNil(t, m)
	assert.Len(t, m, 3)
}
