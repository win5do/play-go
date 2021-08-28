package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestX(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{1, 1, 2, 2, 3}))    // invalid input
	assert.Equal(t, 3, majorityElement([]int{1, 1, 2, 2, 3, 2})) // invalid input

	assert.Equal(t, 2, majorityElement([]int{1, 1, 2, 2, 3, 2, 2}))
	assert.Equal(t, -2, majorityElement([]int{-2}))
	assert.Equal(t, -1, majorityElement([]int{}))
}
