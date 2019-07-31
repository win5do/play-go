package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPath(t *testing.T) {
	matrix := [][]string{
		{"a", "b", "t", "g"},
		{"c", "f", "c", "s"},
		{"j", "d", "e", "h"},
	}

	cs := []struct {
		str      string
		expected bool
	}{
		{str: "bfce", expected: true},
		{str: "bfcehsc", expected: false},
		{str: "bfcehsx", expected: false},
		{str: "bfcehsgt", expected: true},
	}
	for _, c := range cs {
		r := hasPath(matrix, c.str)
		t.Log(r)
		assert.Equal(t, c.expected, r)
	}
}
