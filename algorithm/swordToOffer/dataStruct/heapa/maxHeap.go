package heapa

type MaxHeap struct {
	MinHeap
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.array[i] > h.array[j]
}
