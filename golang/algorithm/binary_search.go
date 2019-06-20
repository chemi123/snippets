package main

import (
	"fmt"
)

type BinarySearchTree interface {
	Add(num int)
	Search(num int) bool
	Traverse()
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Add(num int) {
	if n == nil {
		fmt.Println("root")
		n = createNode(num)
		return
	}

	var prev *Node
	for {
		if n.Val == num {
			return
		}

		prev = n
		if n.Val > num {
			n = n.Left
			if n == nil {
				prev.Left = createNode(num)
				return
			}
			continue
		}

		n = n.Right
		if n == nil {
			prev.Right = createNode(num)
			return
		}
	}
}

func (n *Node) Search(num int) bool {
	for n != nil {
		if n.Val == num {
			return true
		}

		if n.Val > num {
			n = n.Left
			continue
		}

		n = n.Right
	}

	return false
}

func (n *Node) Traverse() {
	if n == nil {
		return
	}
	n.Left.Traverse()
	fmt.Println(n.Val)
	n.Right.Traverse()
}

func createNode(val int) *Node {
	return &Node{Val: val, Left: nil, Right: nil}
}

func createSampleBinarySearchTree() *Node {
	root := createNode(100)
	root.Add(50)
	root.Add(25)
	root.Add(10)
	root.Add(35)
	root.Add(75)
	root.Add(65)
	root.Add(90)
	root.Add(150)
	root.Add(125)
	root.Add(175)

	return root
}

func main() {
	root := createSampleBinarySearchTree()
	root.Add(70)
	root.Traverse()
}
