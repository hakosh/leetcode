package merge_k_sorted_lists

type MinHeap struct {
	heap []int
}

func CreateQueue() MinHeap {
	return MinHeap{[]int{}}
}

func (h MinHeap) Size() int {
	return len(h.heap)
}

func (h MinHeap) Empty() bool {
	return h.Size() == 0
}

func (h MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) swap(i, j int) {
	(*h).heap[i], (*h).heap[j] = (*h).heap[j], (*h).heap[i]
}

func (h MinHeap) compare(i, j int) bool {
	return h.heap[i] < h.heap[j]
}

func (h *MinHeap) Push(val int) {
	(*h).heap = append((*h).heap, val)
	h.siftUp()
}

func (h *MinHeap) siftUp() {
	idx := h.Size() - 1

	for idx > 0 && h.compare(idx, h.parent(idx)) {
		h.swap(idx, h.parent(idx))
		idx = h.parent(idx)
	}
}

func (h *MinHeap) Pop() int {
	v := h.heap[0]

	if h.Size() > 1 {
		h.swap(0, h.Size()-1)
	}

	(*h).heap = (*h).heap[:h.Size()-1]
	h.siftDown()

	return v
}

func (h *MinHeap) siftDown() {
	idx := 0
	l, r := h.leftChild(idx), h.rightChild(idx)

	for (l < h.Size() && h.compare(l, idx)) || (r < h.Size() && h.compare(r, idx)) {
		smaller := l
		if r < h.Size() && h.compare(r, l) {
			smaller = r
		}

		h.swap(smaller, idx)
		idx = smaller
		l, r = h.leftChild(idx), h.rightChild(idx)
	}
}
