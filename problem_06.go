package main

import "fmt"

/**
DISCLAMER : XOR list does not really work  with go's GC
*/

type XORNode struct {
	data int
	both *XORNode
}

func XOR(left, right *XORNode) int {
	return left.data ^ right.data
}

func (node *XORNode) add(data int) {

	newNode := &XORNode{}
	newNode.data = data
	newNode.both = &XORNode{data: XOR(node, nil)}

	if node != nil {
		next := &XORNode{data: XOR(newNode.both, nil)}
		node.both = &XORNode{data: XOR(newNode, next)}
	} else {
		node = newNode
	}
}

func (node *XORNode) get(index int) *XORNode {
	current := node
	prev := &XORNode{}
	next := &XORNode{}

	for ; index > 0 && next != nil; index-- {
		next = &XORNode{data: XOR(prev, current)}
		prev = current
		current = next

		if next == nil && index > 1 {
			fmt.Errorf("Index out of bound!")
		}
	}
	return current
}
