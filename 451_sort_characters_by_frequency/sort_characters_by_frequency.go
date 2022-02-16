package sort_characters_by_frequency

import (
	"container/heap"
	"sort"
	"strings"
)

type MaxHeap [][2]int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i][1] > m[j][1]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.([2]int))
}

func (m *MaxHeap) Pop() interface{} {
	num := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return num
}

func frequencySortHeap(s string) string {
	m := [127]int{}
	for _, char := range s {
		m[char]++
	}

	h := &MaxHeap{}
	for i := 0; i < 127; i++ {
		if m[i] > 0 {
			heap.Push(h, [2]int{i, m[i]})
		}
	}

	sorted := strings.Builder{}

	for h.Len() > 0 {
		char := heap.Pop(h).([2]int)

		chars := make([]byte, char[1])
		for i := 0; i < char[1]; i++ {
			chars[i] = byte(char[0])
		}
		sorted.Write(chars)
	}

	return sorted.String()
}

func frequencySort(s string) string {
	m := [127]int{}

	for _, char := range s {
		m[char]++
	}

	keys := make([]int, 127)
	for i := 0; i < 127; i++ {
		keys[i] = i
	}

	sort.Slice(keys, func(i, j int) bool {
		a := m[keys[i]]
		b := m[keys[j]]

		return a > b
	})

	sorted := strings.Builder{}

	for _, key := range keys {
		times := m[key]

		chars := make([]byte, times)
		for i := 0; i < times; i++ {
			chars[i] = byte(key)
		}
		sorted.Write(chars)
	}

	return sorted.String()
}
