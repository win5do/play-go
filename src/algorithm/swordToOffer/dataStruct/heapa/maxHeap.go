package heapa

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Pop() interface{} {
	l := len(h.array)
	e := h.array[l-1]
	h.array = h.array[:l-1]
	return e
}

func (h *MaxHeap) Push(e interface{}) {
	h.array = append(h.array, e.(int))
}

func (h *MaxHeap) Len() int {
	return len(h.array)
}

func (h *MaxHeap) Less(i, j int) bool {
	return i > j
}

func (h *MaxHeap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap) Top() int {
	return h.array[0]
}

func (h *MaxHeap) Value() []int {
	return h.array
}
