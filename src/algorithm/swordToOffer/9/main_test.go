package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	que := new(Queue)
	que.Push(1)
	que.Push(2)
	que.Push(3)
	t.Log(que.Pop())
	t.Log(que.Pop())
	t.Log(que.Pop())
	que.Push(4)
	t.Log(que.Pop())
	t.Log(que.Pop())
}
