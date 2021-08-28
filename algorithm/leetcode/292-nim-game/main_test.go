package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim_dp(t *testing.T) {
	for i := 1; i <= 100; i++ {
		assert.Equal(t, canWinNim_bitwise(i), canWinNim_dp(i))
	}
}
