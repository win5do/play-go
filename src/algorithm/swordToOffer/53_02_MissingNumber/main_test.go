package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	cs := []struct {
		arr      []int
		expected int
	}{
		{arr: []int{1, 2}, expected: 0},
		{arr: []int{0, 1}, expected: 2},
		{arr: []int{0, 2}, expected: 1},
	}

	for i, v := range cs {
		r := getMissingNumber(v.arr)
		assert.Equal(t, v.expected, r, i)
	}
}
