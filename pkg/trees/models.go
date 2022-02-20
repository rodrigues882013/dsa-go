package trees

type (
	Node struct {
		Value int
		Left *Node
		Right *Node
	}

	Tree interface {
		Search(n *Node, value int) bool
		Insert(value int)
	}
)