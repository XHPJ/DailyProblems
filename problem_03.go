package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) insert(value int) *Tree {
	if tree.root == nil {
		tree.root = &Node{val: value, left: nil, right: nil}
	} else {
		tree.root.insert(value)
	}
	return tree
}

func (node *Node) insert(value int) {
	if node == nil {
		return
	}
	if value > node.val {
		if node.right == nil {
			node.right = &Node{val: value, left: nil, right: nil}
		} else {
			node.right.insert(value)
		}
	} else {
		if node.left == nil {
			node.left = &Node{val: value, left: nil, right: nil}
		} else {
			node.left.insert(value)
		}
	}
}

func (tree *Tree) printTree() {
	tree.root.printNode("m")
}

func (node *Node) printNode(side string) {
	if node == nil {
		fmt.Println("[]")
	} else {
		fmt.Print(side, ":", node.val, "  ")
		node.left.printNode("l")
		node.right.printNode("r")
	}
}

// func main() {

// 	tree := &Tree{}
// 	tree.printTree()
// 	tree.insert(3)
// 	tree.insert(2)
// 	tree.insert(5)
// 	tree.insert(4)
// 	tree.insert(1)
// 	tree.insert(6)

// 	fmt.Println(tree.root.left)
// 	tree.printTree()
// }
