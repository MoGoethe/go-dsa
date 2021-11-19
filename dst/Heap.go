package dst

import "fmt"

// 堆：特殊的完全二叉树
// 最小堆：父节点的值比每一个子节点的值都要小
// 最大堆：父节点的值比每一个子节点的值都要大
// MaxHeap strtct
type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.maxHeapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) Extract() int {
	extracted := h.items[0]
	l := len(h.items) - 1

	if len(h.items) == 0 {
		fmt.Println("cannot extract because heap length is 0")
		return -1
	}
	h.items[0] = h.items[l]
	h.items = h.items[:l]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.items) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.items[l] > h.items[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int {
	return 2*i + 1
}

// Get the reight child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}
