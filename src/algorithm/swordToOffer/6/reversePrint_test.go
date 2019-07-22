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

func TestReverseLinkedList(t *testing.T) {
	list := reverseLinkedList(makeList(10))
	PrintList(list)
}
