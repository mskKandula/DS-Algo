package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(val int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{
			data: val,
		}
	} else {
		t.root.insert(val)
	}
	return t
}

func (n *Node) insert(val int) {
	if n == nil {
		return
	} else if val <= n.data {
		if n.left == nil {
			n.left = &Node{data: val}
		} else {
			n.left.insert(val)
		}

	} else {
		if n.right == nil {
			n.right = &Node{data: val}
		} else {
			n.right.insert(val)
		}
	}

}

func (t *BinaryTree) PostOrder(n *Node) {
	if n == nil {
		return
	}
	t.PostOrder(n.left)
	t.PostOrder(n.right)
	fmt.Printf("%d -> ", n.data)
}

func main() {
	fmt.Println("Depth First Search")
	bTree := &BinaryTree{}

	bTree.insert(5).
		insert(6).
		insert(7).
		insert(8).
		insert(9).
		insert(4).
		insert(3).
		insert(2).
		insert(1)

	bTree.PostOrder(bTree.root)
}
