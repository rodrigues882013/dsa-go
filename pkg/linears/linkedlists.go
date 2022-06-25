package linears

type Node struct {
	Val TreeNode
	Key int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList {
		Head: &Node{},
		Tail: &Node{},
	}
}


type Operation interface {
	Append(node *Node)
	AppendLeft(node *Node)
	Pop() *Node
	PopLeft() *Node
}

func (l *LinkedList) Append(node *Node)  {

	if l.Head.Next == nil {
		node.Next = l.Tail
		node.Prev = l.Head
		l.Head.Next = node
		l.Tail.Prev = node

	} else {
		node.Prev = l.Tail.Prev
		l.Tail.Prev.Next = node
		l.Tail.Prev = node
	}
	l.Size += 1
}

func (l* LinkedList) AppendLeft(node *Node)  {
	if l.Head.Next == nil {
		node.Next = l.Tail
		node.Prev = l.Head
		l.Head.Next = node
		l.Tail.Prev = node

	} else {
		node.Next = l.Head.Next
		l.Head.Next.Prev = node
		l.Head.Next = node
	}
	l.Size += 1
}


func (l *LinkedList) Pop() *Node  {
	popped := l.Tail.Prev
	l.Tail.Prev = l.Tail.Prev.Prev
	l.Tail.Prev.Next = l.Tail
	l.Size -= 1
	return popped
}

func (l *LinkedList) PopLeft() *Node {
	popped := l.Head.Next
	l.Head.Next = l.Head.Next.Next
	l.Head.Next.Prev = l.Head
	l.Size -= 1
	return popped
}

