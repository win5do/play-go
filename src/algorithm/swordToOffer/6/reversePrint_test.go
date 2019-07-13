package main

import (
	"testing"
)

func TestReversePrintByStack(t *testing.T) {
	reverseByStack(makeList(10))
}

func TestReverseByRecursion(t *testing.T) {
	reverseByRecursion(makeList(10))
}
