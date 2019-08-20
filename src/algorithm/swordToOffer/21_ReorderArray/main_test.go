package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderArray(t *testing.T) {
	cs := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 2, 1, 2, 1, 2},
			expected: []int{1, 1, 1, 2, 2, 2},
		},
		{
			arr:      []int{1, 1, 1, 1, 2, 1},
			expected: []int{1, 1, 1, 1, 1, 2},
		},
		{
			arr:      []int{1, 1, 1, 1, 1, 2},
			expected: []int{1, 1, 1, 1, 1, 2},
		},
		{
			arr:      []int{2, 1, 1, 1, 1, 2},
			expected: []int{1, 1, 1, 1, 2, 2},
		},
		{
			arr:      []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			arr:      []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
	}

	for i, v := range cs {
		ReorderArray(v.arr)
		assert.Equal(t, v.expected, v.arr, i)
	}
}
