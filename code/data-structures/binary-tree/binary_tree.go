// Binary Tree implementation in Go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree defines the structure of the binary tree.
// Each node can be connected to one or at the most
// two nodes.
//
// First node is called as "root" and the nodes with
// no children are called as "leaves"
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// traverse reveals how to visit all nodes
func traverse(t *Tree) {
	if t == nil {
		return
	}
	// using recursion to visit all the nodes
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

// create is just used to seed the tree with random numbers
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	// check if the tree is empty
	if t == nil {
		// will set the node as root
		return &Tree{nil, v, nil}
	}
	// check if the value already exists
	if v == t.Value {
		// return without inserting
		return t
	}
	// check if values has to be inserted in left or right node
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	// create a tree with random values
	tree := create(10)
	fmt.Println("The value of the root of tree is:", tree.Value)
	// traverse the tree
	traverse(tree)
	fmt.Println()
	// insert values
	tree = insert(tree, -10)
	tree = insert(tree, -2)

	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of tree is:", tree.Value)
}
