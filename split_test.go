package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}
	chunks := Split(arr, 3)
	assert.Equal(t, []string{"1", "2", "3"}, chunks[0])
	assert.Equal(t, []string{"4", "5", "6"}, chunks[1])
	assert.Equal(t, []string{"7", "8", "9"}, chunks[2])
	assert.Equal(t, []string{"10", "11"}, chunks[3])
}

// 11805250	       109 ns/op	      96 B/op	       1 allocs/op
func BenchmarkSplit(b *testing.B) {
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

	for i := 0; i < b.N; i++ {
		Split(arr, 3)
	}
}
