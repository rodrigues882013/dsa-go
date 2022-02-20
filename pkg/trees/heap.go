package trees

import "fmt"

type Heap struct {
	Storage []int
	Size int `default:"0"`
}

func (h *Heap) bubbleUp(lastIndex int) {
	parentIndex := lastIndex / 2

	if h.Storage[lastIndex] > h.Storage[parentIndex] {
		tmp := h.Storage[lastIndex]
		h.Storage[lastIndex] = h.Storage[parentIndex]
		h.Storage[parentIndex] = tmp
		h.bubbleUp(parentIndex)
	}
	var m = map[int]int{}
	fmt.Println(m[1])
}

func (h *Heap) bubbleDown(firstIndex int) {

	// Just two elements in list generate a special treatment
	if h.Size == 2 {
		tmp := h.Storage[1]

		if tmp > h.Storage[0] {
			h.Storage[1] = h.Storage[0]
			h.Storage[0] = tmp
		}

		return
	}

	leftChild := firstIndex * 2
	rightChild := (firstIndex * 2) + 1

	if leftChild > h.Size - 2 || rightChild > h.Size - 2 {
		return
	}

	if h.Storage[firstIndex] > h.Storage[leftChild + 1] && h.Storage[firstIndex] > h.Storage[rightChild + 1] {
		return
	}

	if h.Storage[leftChild + 1] >= h.Storage[rightChild + 1] {
		tmp := h.Storage[firstIndex]
		h.Storage[firstIndex] = h.Storage[leftChild + 1]
		h.Storage[leftChild + 1] = tmp
		h.bubbleDown(leftChild + 1)

	} else if h.Storage[leftChild + 1] < h.Storage[rightChild + 1] {
		tmp := h.Storage[firstIndex]
		h.Storage[firstIndex] = h.Storage[rightChild + 1]
		h.Storage[rightChild + 1] = tmp
		h.bubbleDown(rightChild + 1)
	}
}

func (h *Heap) Search(n *Node, value int) bool {
	panic("implement me")
}

func (h *Heap) getMin() int {
	return h.Storage[0]
}

func (h *Heap) GetMax() int {
	return h.Storage[0]
}

func (h *Heap) Insert(value int) {
	h.Storage = append(h.Storage, value)
	h.Size += 1

	if len(h.Storage) > 1 {
		h.bubbleUp(h.Size - 1)
	}
}

func (h *Heap) Pop() {
	h.Storage[0] = h.Storage[h.Size - 1]
	h.Size = h.Size - 1

	if h.Size == 0 {
		h.Storage = nil
	}

	h.bubbleDown(0)
}




