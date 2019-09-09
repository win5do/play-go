package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	cs := []struct {
		arr      []int
		k        int
		expected int
	}{
		{arr: []int{1, 2, 3, 3, 3, 3, 4, 5}, k: 3, expected: 4},
		{arr: []int{1, 2, 3, 4}, k: 1, expected: 1},
		{arr: []int{1, 2, 3, 4}, k: 4, expected: 1},
		{arr: []int{1, 2, 3, 4}, k: 2, expected: 1},
		{arr: []int{}, k: 2, expected: 0},
		{arr: []int{1, 1, 2}, k: 1, expected: 2},
		{arr: []int{1, 1}, k: 2, expected: 0},
		{arr: []int{1, 2, 2}, k: 2, expected: 2},
	}

	for i, v := range cs {
		r := getNumberOfK(v.arr, v.k)
		assert.Equal(t, v.expected, r, i)
	}
}
