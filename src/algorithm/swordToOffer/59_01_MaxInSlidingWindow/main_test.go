package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	cs := []struct {
		arr      []int
		size     int
		expected []int
	}{
		{arr: []int{2, 3, 4, 2, 6, 2, 5, 1}, size: 3, expected: []int{4, 4, 6, 6, 6, 5}},
		{arr: []int{1, 2}, size: 3, expected: []int{2}},
		{arr: nil, size: 3, expected: []int{}},
		{arr: []int{1, 2, 3, 4}, size: 0, expected: []int{}},
		{arr: []int{1, 2, 3, 4, 5}, size: 2, expected: []int{2, 3, 4, 5}},
		{arr: []int{5, 4, 3, 2, 1}, size: 2, expected: []int{5, 4, 3, 2}},
	}

	for i, v := range cs {
		r := maxInWindows(v.arr, v.size)
		assert.Equal(t, v.expected, r, i)
	}
}
