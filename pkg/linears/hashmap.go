package linears

type HashMap struct {
	Storage []*LinkedList
}

func hashing(key int) int {
	return key % 720
}

func NewHashMap() *HashMap {
	return &HashMap{make([]*LinkedList, 720) }
}

func (h *HashMap) Put(key, value int)  {
	keyHashed := hashing(key)

	if h.Storage[keyHashed] == nil {
		h.Storage[keyHashed] = NewLinkedList()
	}
	h.Storage[keyHashed].Append(&Node{Key: key, Val: value})
}

func (h *HashMap) Get(key int) int {
	keyHashed := hashing(key)
	currentList := h.Storage[keyHashed]

	if currentList == nil {
		panic("No such key")
	}

	runner := currentList.Head
	for runner.Next != nil && runner.Key != key {
		runner = runner.Next
	}

	if runner == nil {
		panic("No such key")
	}

	return runner.Val
}
