package linears

type Stack struct {
	storage *LinkedList
}

func NewStack() *Stack {
	return &Stack{NewLinkedList() }
}

func (s *Stack) Push(value int)  {
	node := &Node{Val: value}
	s.storage.Append(node)
}

func (s *Stack) Pop() int  {
	if s.storage.Size == 0 {
		panic("Trying popping from a empty stack")
	}
	popped := s.storage.Pop()
	return popped.Val
}

func (s *Stack) Peek() int {
	if s.storage.Size == 0 {
		panic("Trying see the peek from a empty stack")
	}

	return s.storage.Tail.Prev.Val
}


