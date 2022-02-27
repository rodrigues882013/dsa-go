package trees

type (

	// Node is a struct holding a node used in AVL, BST and RBL initially and considering a int as value.
	Node struct {
		Value int
		Left *Node
		Right *Node
		Color bool
	}

	// Tree is an interface holding the most commons operations regardless in a any kind of tree
	Tree interface {
		Search(n *Node, value int) bool
		Insert(value int)
	}
)