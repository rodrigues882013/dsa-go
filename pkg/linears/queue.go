package linears

type Queue struct {
	storage *LinkedList
}

func NewQueue() *Queue {
	return &Queue{NewLinkedList() }
}

func (q *Queue) Enqueue(value int)  {
	node := &Node{Val: value}
	q.storage.Append(node)
}

func (q *Queue) Dequeue() int  {
	if q.storage.Size == 0 {
		panic("Trying pop from a empty stack")
	}
	dequeued := q.storage.PopLeft()
	return dequeued.Val
}

func (q *Queue) Rear() int {
	if q.storage.Size == 0 {
		panic("Trying see the rear from a empty stack")
	}

	return q.storage.Tail.Prev.Val
}

func (q *Queue) Front() int {
	if q.storage.Size == 0 {
		panic("Trying see the front from a empty stack")
	}

	return q.storage.Head.Next.Val
}
