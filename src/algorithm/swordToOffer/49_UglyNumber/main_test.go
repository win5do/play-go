package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	for i := 0; i <= 100; i++ {
		r1 := getUglyNumber(i)
		r2 := getUglyNumber2(i)
		assert.Equal(t, r1, r2, i)
		t.Log(r1)
	}
}
