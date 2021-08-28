package heapa_test

import (
	"container/heap"
	"fmt"
	"testing"

	"playGo/algorithm/swordToOffer/dataStruct/heapa"
)

func TestMinHeap(t *testing.T) {
	h := new(heapa.MinHeap)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 5)

	fmt.Println(h.Value())
	fmt.Println(heap.Pop(h))
	heap.Push(h, 0)
	fmt.Println(h.Value())
}

func TestMaxHeap(t *testing.T) {
	h := new(heapa.MaxHeap)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 5)

	fmt.Println(h.Value())
	fmt.Println(heap.Pop(h))
	heap.Push(h, 0)
	fmt.Println(h.Value())
}
