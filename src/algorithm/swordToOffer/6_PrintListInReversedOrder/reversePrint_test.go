package main

import (
	"testing"
)

func TestReversePrintByStack(t *testing.T) {
	reversePrint_stack(makeList(10))
}

func TestReverseByRecursion(t *testing.T) {
	reversePrint_recursion(makeList(10))
}

func TestReverseList(t *testing.T) {
	list := reverseList(makeList(10))
	PrintList(list)
}
