package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	bst := BST{}
	nodeSize := 0

	fmt.Scan(&nodeSize)
	for i := 0; i < nodeSize; i++ {
		bst.Insert(getInput())
	}

	size := reflect.TypeOf(bst.root).Size()
	fmt.Println(size)

}

func (bst *BST) Insert(val int) {
	bst.InsertRec(bst.root, val)
}

func (bst *BST) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	}
	if val > node.data {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}

func getInput() int {
	var userInput int
	fmt.Scan(&userInput)
	return userInput
}
