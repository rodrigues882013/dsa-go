package trees

import (
	"os"
)

func insert(n *Node, value int, isAvl bool) *Node {
	if n == nil {
		return &Node{
			Value: value,
		}
	}

	if value > n.Value {

		if n.Right == nil {
			n.Right = &Node{Value: value}

		} else {
			n.Right = insert(n.Right, value, isAvl)
		}

	} else if value < n.Value {

		if n.Left == nil {
			n.Left = &Node{Value: value}

		} else {
			n.Left = insert(n.Left, value, isAvl)
		}

	} else {
		os.Exit(0)
	}

	if isAvl {
		return rebalance(n, value)
	}

	return n
}

func search(n *Node, value int) bool {

	if n == nil {
		return false
	}

	if value > n.Value {
		return search(n.Right, value)
	}

	if value < n.Value {
		return search(n.Left, value)
	}

	return true

}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b

}

func getHeight(node *Node) int {
	if node == nil {
		return 0

	} else if node.Left == nil && node.Right == nil {
		return 1

	} else if node.Left == nil {
		return 1 + getHeight(node.Right)

	} else if node.Right == nil {
		return 1 + getHeight(node.Left)

	} else {
		return 1 + max(getHeight(node.Left), getHeight(node.Right))
	}
}

func leftRotate(node *Node) *Node {

	subtree := node.Right
	tmp := subtree.Left

	subtree.Left = node
	node.Right = tmp

	return subtree
}

func rightRotate(node *Node) *Node {
	subtree := node.Left
	tmp := subtree.Right

	subtree.Right = node
	node.Left = tmp

	return subtree
}


func rebalance(node *Node, data int) *Node {
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	balance := leftHeight - rightHeight

	if balance < -1 {
		if data < node.Right.Value {

			// Right-Left
			node.Right = rightRotate(node.Right)
			node = rightRotate(node)

		} else {

			// Left
			node = leftRotate(node)
		}
	} else if balance > 1 {
		if data > node.Left.Value {

			// Left-Right
			node.Left = leftRotate(node.Left)
			node = leftRotate(node)

		} else {
			// Right
			node = rightRotate(node)
		}
	}

	return node
}


func Search(n *Node, value int) bool {
	if n == nil {
		return false
	}

	if n.Value == value {
		return true

	} else {
		return search(n, value)
	}
}
