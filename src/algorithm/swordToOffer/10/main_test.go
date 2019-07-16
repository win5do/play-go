package main

import (
	"testing"
)

func TestFibonacciSlow(t *testing.T) {
	t.Log(fibonacciSlow(0))
	t.Log(fibonacciSlow(10))
	t.Log(fibonacciSlow(20))
}

func TestFibonacciLoop(t *testing.T) {
	t.Log(fibonacciLoop(0))
	t.Log(fibonacciLoop(10))
	t.Log(fibonacciLoop(20))
	t.Log(fibonacciLoop(50))
	t.Log(fibonacciLoop(100))
}

func TestFibonacciRecurse(t *testing.T) {
	t.Log(fibonacciRecurse(0))
	t.Log(fibonacciRecurse(1))
	t.Log(fibonacciRecurse(10))
	t.Log(fibonacciRecurse(20))
	t.Log(fibonacciRecurse(50))
	t.Log(fibonacciRecurse(100))
}
