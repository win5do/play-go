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
		{arr: []int{1, 2, 1, 2, 2}, expected: 2},
		{arr: []int{}, expected: 0},
		{arr: []int{1, 2, 1, 2, 3}, expected: 0},
		{arr: []int{1, 1, 2, 2}, expected: 1},
		{arr: []int{1, 2, 1, 2}, expected: 1},
	}

	for i, v := range cs {
		// 有两个解的两个函数不一致
		r1 := moreThanHalfNum_partition(v.arr)
		r2 := moreThanHalfNum_count(v.arr)
		assert.Equal(t, v.expected, r1, i)
		assert.Equal(t, v.expected, r2, i)
	}
}
