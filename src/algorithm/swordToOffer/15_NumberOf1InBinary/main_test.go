package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOf1_unsign(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(NumberOf1_unsign(i))
	}
}

func TestNumberOf1(t *testing.T) {
	for i := -100; i < 100; i++ {
		r1 := NumberOf1_leftShift(i)
		r2 := NumberOf1_and(i)
		t.Log(r1, r2)
		assert.Equal(t, r1, r2)
	}
}
