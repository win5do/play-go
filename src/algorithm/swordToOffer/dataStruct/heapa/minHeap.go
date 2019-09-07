package heapa

type MinHeap struct {
	array []int
}

func (h *MinHeap) Pop() interface{} {
	l := len(h.array)
	e := h.array[l-1]
	h.array = h.array[:l-1]
	return e
}

func (h *MinHeap) Push(e interface{}) {
	h.array = append(h.array, e.(int))
}

func (h *MinHeap) Len() int {
	return len(h.array)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.array[i] < h.array[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MinHeap) Top() int {
	return h.array[0]
}

func (h *MinHeap) Bottom() int {
	l := len(h.array)
	return h.array[l-1]
}

func (h *MinHeap) Value() []int {
	return h.array
}
