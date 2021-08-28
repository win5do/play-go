package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumeric(t *testing.T) {
	cs := []struct {
		str      string
		expected bool
	}{
		{str: "100", expected: true},
		{str: "+100", expected: true},
		{str: "-100", expected: true},
		{str: "+-100", expected: false},
		{str: "12.34e56", expected: true},
		{str: "12.34e+56", expected: true},
		{str: "+12.34e+56", expected: true},
		{str: "-12.34e-56", expected: true},
		{str: "-12.34E+56", expected: true},
		{str: "-12.34e-56.", expected: false},
		{str: "12.", expected: false},
		{str: "12e", expected: false},
		{str: "12.34e", expected: false},
		{str: "12.e34", expected: false},
		{str: "1+2.12e34", expected: false},
		{str: "", expected: false},
	}

	for i, v := range cs {
		r := isNumeric(v.str)
		assert.Equal(t, v.expected, r, i)
	}
}
