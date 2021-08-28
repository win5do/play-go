package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	cs := []struct {
		str      string
		pattern  string
		expected bool
	}{
		{str: "", pattern: "", expected: true},
		{str: "abc", pattern: "ax*bc", expected: true},
		{str: "abc", pattern: "...", expected: true},
		{str: "abc", pattern: "abcd*e*", expected: true},
		{str: "abc", pattern: "abcde", expected: false},
		{str: "abcde", pattern: "abc", expected: false},
		{str: "abcde", pattern: ".*", expected: true},
	}

	for _, v := range cs {
		r := match(v.str, v.pattern)
		assert.Equal(t, v.expected, r)
	}
}
