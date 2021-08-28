package chanQuene

import (
	"testing"
)

func TestQuene_Size(t *testing.T) {
	que := NewQuene()
	t.Log(que.Size())
	que.Push(1)
	t.Log(que.Size())
	que.Pop()
	t.Log(que.Size())
}
