package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := []string{"2", "1", "4", "3", "5"}
	got := diff(a, b)
	assert.Equal(t, got, []string{"4", "5"})
}

// 8266914               142 ns/op              80 B/op          1 allocs/op
func BenchmarkDiff(b *testing.B) {
	arr1 := []string{"1", "2", "3"}
	arr2 := []string{"2", "1", "4", "3", "5"}
	for i := 0; i < b.N; i++ {
		diff(arr1, arr2)
	}
}

func TestDiffmulti(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := []string{"2", "1", "4", "3", "5"}
	got := diffmulti(a, b)
	if !reflect.DeepEqual(got, []string{"4", "5"}) && !reflect.DeepEqual(got, []string{"5", "4"}) {
		// reverse also ok
		t.Error(got)
	}
}

func BenchmarkDiffmulti(b *testing.B) {
	arr1 := []string{"1", "2", "3"}
	arr2 := []string{"2", "1", "4", "3", "5"}
	for i := 0; i < b.N; i++ {
		diffmulti(arr1, arr2)
	}
}
