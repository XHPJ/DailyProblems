package main

import "fmt"

/*
Task: 		A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
			Given the root to a binary tree, count the number of unival subtrees.

			For example, the following tree has 5 unival subtrees:
				 0
				/ \
			   1   0
				  / \
				 1   0
				/ \
			   1   1
**/

//Basic tree structure without ordering and allow for incompletenes
type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

//NOT THE BEST WAY TO INSERT A NODE BUT THIS IS NOT THE SCOPE OF THIS TASK!
func (tree *BinaryTree) SetLeft(value int) *BinaryTree {
	tree.left = &BinaryTree{value: value, left: nil, right: nil}
	return tree
}

func (tree *BinaryTree) SetRight(value int) *BinaryTree {
	tree.right = &BinaryTree{value: value, left: nil, right: nil}
	return tree
}

func CountUivalTrees(tree *BinaryTree) int {
	//In case we reached the last node return the acc
	if tree == nil {
		return 0
	}
	if tree.left == nil && tree.right == nil {
		return 1
	}

	count := 0
	//in case of matching values the amount of unival trees equals the sum of the acc
	if tree.left != nil && tree.right == nil {
		if tree.left.value == tree.value {
			count += 1 + CountUivalTrees(tree.left)
		}
	}
	if tree.right != nil && tree.left == nil {
		if tree.right.value == tree.value {
			count += 1 + CountUivalTrees(tree.right)
		}
	}
	if tree.right != nil && tree.left != nil {
		count += CountUivalTrees(tree.right) + CountUivalTrees(tree.left)
	}
	if tree.value == tree.left.value && tree.left.value == tree.value {
		count++
	}
	return count
}

func main() {
	//same tree as in the example above
	tree := BinaryTree{value: 0, left: nil, right: nil}
	tree.SetLeft(1)
	tree.SetRight(0)
	tree.right.SetLeft(1)
	tree.right.SetRight(0)
	tree.right.left.SetLeft(1)
	tree.right.left.SetRight(1)

	fmt.Println(CountUivalTrees(&tree))
}
