package trees

// AVLTree is a struct that represent a balanced binary tree
type AVLTree struct {
	Root *Node
}

func (a *AVLTree) Search(n *Node, value int) bool {
	return Search(n, value)
}

func (a *AVLTree) Insert(value int) {
	if Search(a.Root, value) {
		return
	}
	a.Root = insert(a.Root, value, true)
}

