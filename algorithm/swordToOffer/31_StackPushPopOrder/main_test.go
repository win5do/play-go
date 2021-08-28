package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPopOrder(t *testing.T) {
	cs := []struct {
		pushOrder []int
		popOrder  []int
		expected  bool
	}{
		{pushOrder: []int{1, 2, 3, 4, 5}, popOrder: []int{4, 5, 3, 2, 1}, expected: true},
		{pushOrder: []int{1, 2, 3, 4, 5}, popOrder: []int{4, 3, 5, 1, 2}, expected: false},
		{pushOrder: []int{1, 2, 3, 4, 5}, popOrder: []int{5, 1, 2, 3, 4}, expected: false},
		{pushOrder: []int{}, popOrder: []int{}, expected: false},
		{pushOrder: []int{1}, popOrder: []int{1, 2}, expected: false},
		{pushOrder: []int{1, 1, 1, 1, 1}, popOrder: []int{1, 1, 1, 1, 1}, expected: true},
		{pushOrder: []int{1, 1, 1, 2, 2}, popOrder: []int{1, 2, 2, 1, 1}, expected: true},
		{pushOrder: []int{1, 1, 1, 2, 2}, popOrder: []int{2, 2, 2, 1, 1}, expected: false},
		{pushOrder: []int{1, 1, 1, 2, 2}, popOrder: []int{1, 2, 1, 2, 1}, expected: true},
		{pushOrder: []int{1, 1, 1, 2, 2}, popOrder: []int{2, 1, 1, 1, 2}, expected: true},
	}

	for i, v := range cs {
		r := isPopOrder(v.pushOrder, v.popOrder)
		assert.Equal(t, v.expected, r, i)
	}
}
