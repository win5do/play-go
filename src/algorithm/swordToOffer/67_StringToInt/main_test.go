package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	cs := []struct {
		str      string
		expected int
		err      error
	}{
		{str: "1234", expected: 1234, err: nil},
		{str: "+1234", expected: 1234, err: nil},
		{str: "-1234", expected: -1234, err: nil},
		{str: "-", expected: 0, err: errInvalidInput},
		{str: "+", expected: 0, err: errInvalidInput},
		{str: "0+-", expected: 0, err: errInvalidInput},
	}

	for i, v := range cs {
		r, err := atoi(v.str)
		assert.Equal(t, v.expected, r, i)
		assert.Equal(t, v.err, err, i)
	}
}
