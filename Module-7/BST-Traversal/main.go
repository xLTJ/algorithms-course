package main

import "fmt"

type Node struct {
	Key    int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BTree struct {
	Root *Node
}

func (t *BTree) Insert(node *Node) {
	var y *Node = nil
	x := t.Root
	for x != nil {
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y
	if y == nil {
		t.Root = node
	} else if node.Key < y.Key {
		y.Left = node
	} else {
		y.Right = node
	}
}

func (t *BTree) PreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.Key)
		t.PreOrder(node.Left)
		t.PreOrder(node.Right)
	}
}

func (t *BTree) PostOrder(node *Node) {
	if node != nil {
		t.PostOrder(node.Left)
		t.PostOrder(node.Right)
		fmt.Println(node.Key)
	}
}

func main() {
	tree := &BTree{}
	keys := []int{10, 4, 17, 1, 4, 16, 21}

	for _, key := range keys {
		tree.Insert(&Node{Key: key})
	}

	fmt.Println("PreOrder:")
	tree.PreOrder(tree.Root)

	fmt.Println("\nPostOrder:")
	tree.PostOrder(tree.Root)
}
