package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProductAfterCutting(t *testing.T) {
	for i := 0; i < 20; i++ {
		r1 := maxProductAfterCutting_dynamic(i)
		r2 := maxProductAfterCutting_greedy(i)
		assert.Equal(t, r1, r2)
		fmt.Printf("n = %v, r = %v\n", i, r1)
	}
}
