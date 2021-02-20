package intslice

import "fmt"

// Sort2 func use binary tree sort
func Sort2(slice []int) []int {
	size := len(slice)
	if size < 2 {
		return slice
	}
	// fmt.Println(slice)

	tree := &BinaryTree{}
	for i := 0; i < size; i++ {
		tree.insert(slice[i])
	}
	tree.sortAsc(slice)
	return slice
}

// Node of binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// BinaryTree for sort
type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) sortAsc(slice []int) {

	// t.root.print()
	t.root.addTo(slice, 0)
}

func (n *Node) addTo(slice []int, index int) int {
	if n.left != nil {
		index = n.left.addTo(slice, index)
	}
	slice[index] = n.data
	index++
	if n.right != nil {
		index = n.right.addTo(slice, index)
	}
	return index
}

func (n *Node) print() {
	if n.left != nil {
		n.left.print()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.print()
	}
}
