package trees

// RBLTree is a struct that represent a red-black tree
type RBLTree struct {
	Root *Node
}

func (a *RBLTree) Search(n *Node, value int) bool {
	return Search(n, value)
}

func (a *RBLTree) Insert(value int) {
	if Search(a.Root, value) {
		return
	}
	a.Root = insert(a.Root, value, true)
}
