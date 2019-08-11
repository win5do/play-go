package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	cs := []struct {
		arr []int
		min int
	}{
		{
			arr: []int{1},
			min: 1,
		},
		{
			arr: []int{1, 2},
			min: 1,
		},
		{
			arr: []int{1, 2, 3, 4, 5, 6},
			min: 1,
		},
		{
			arr: []int{2, 1},
			min: 1,
		},
		{
			arr: []int{3, 4, 5, 5, 1, 1, 2, 3},
			min: 1,
		},
	}

	for _, c := range cs {
		r := min(c.arr)
		assert.Equal(t, c.min, r)
	}
}
