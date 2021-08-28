package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	cs := []struct {
		seq      []int
		expected bool
	}{
		{seq: []int{5, 7, 6, 9, 11, 10, 8}, expected: true},
		{seq: []int{1, 2, 3}, expected: true},
		{seq: []int{2, 1, 3}, expected: true},
		{seq: []int{3, 1, 2}, expected: false},
		{seq: []int{1}, expected: true},
		{seq: []int{1, 2}, expected: true},
		{seq: []int{2, 1}, expected: true},
	}

	for i, v := range cs {
		r := verifySequenceOfBST(v.seq)
		assert.Equal(t, v.expected, r, i)
	}
}
