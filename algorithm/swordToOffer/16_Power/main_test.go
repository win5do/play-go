package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	cs := []struct {
		base     float64
		exponent int
	}{
		{base: 2, exponent: 8},
		{base: 2, exponent: -8},
		{base: 2, exponent: 0},
		{base: 2, exponent: 1},
		{base: 2, exponent: -1},
		{base: 2.5, exponent: 8},
		{base: 2.5, exponent: -8},
	}

	for _, c := range cs {
		r1 := math.Pow(c.base, float64(c.exponent))
		r2 := power_loop(c.base, c.exponent)
		r3 := power_recurse(c.base, c.exponent)
		t.Log(r1)
		assert.Equal(t, r1, r2)
		assert.Equal(t, r1, r3)
	}
}
