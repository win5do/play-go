package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackWithMin(t *testing.T) {
	stack := new(stackWithMin)
	stack.Push(3)
	require.Equal(t, 3, stack.Min())

	stack.Push(4)
	require.Equal(t, 3, stack.Min())

	stack.Push(3)
	require.Equal(t, 3, stack.Min())

	stack.Push(2)
	require.Equal(t, 2, stack.Min())

	stack.Pop()
	require.Equal(t, 3, stack.Min())
}
