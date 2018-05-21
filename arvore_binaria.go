package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

type tree struct {
	root *node
}

func main() {
	var t tree
	t.insert(5)
	t.insert(6)
	t.inOrder(t.root)
}

func (root *node) insertAux(leaf *node){
    if leaf.value < root.value {
        if root.right == nil {
            root.right = leaf
            return
        }
		root.left.insertAux(leaf)
	} else {
	    if root.left == nil {
	        root.left = leaf
	        return
	    }
		root.right.insertAux(leaf)
	}
}

func (t *tree) insert(v int) {
    if t.root == nil {
		t.root = &node{v, nil, nil}
		return
	}
	t.root.insertAux(&node{v, nil, nil})
}

func (t *tree) inOrder(n *node) {
	if n.left != nil {
		t.inOrder(n.left)
	}
	fmt.Printf("[%v]", n.value)
	if n.right != nil {
		t.inOrder(n.right)
	}
	fmt.Println()
}
