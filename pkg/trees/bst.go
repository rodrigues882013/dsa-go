package trees

// BSTree is a struct that represent a unbalanced binary tree
type BSTree struct {
	Root *Node
}

func (b *BSTree) Search(n *Node, value int) bool {
	return Search(n, value)
}

func (b *BSTree) Insert(value int) {
	b.Root = insert(b.Root, value, false)
}


