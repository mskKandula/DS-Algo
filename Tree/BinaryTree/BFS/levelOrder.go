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

type Queue [](*Node)

func (self *Queue) Enqueue(val *Node) {
	*self = append(*self, val)
}

func (self *Queue) Dequeue() *Node {
	if self.IsEmpty() {
		return nil
	}
	val := (*self)[0]

	*self = (*self)[1:]

	return val

}

func (self *Queue) IsEmpty() bool {
	return len(*self) == 0
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

func (t *BinaryTree) LevelOrder(n *Node) {
	if n == nil {
		return
	}

	queue := &Queue{}

	queue.Enqueue(n)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		fmt.Print(node.data)
		if node.left != nil {
			queue.Enqueue(node.left)
		}

		if node.right != nil {
			queue.Enqueue(node.right)
		}

	}
	return
}

func main() {
	fmt.Println("Breadth First Search")
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

	bTree.LevelOrder(bTree.root)
}
