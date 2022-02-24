package tree

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

func createNode(value int) *Node {
	return &Node{value: value}
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.left.Traverse()
	node.Print()
	node.right.Traverse()

}

//Camls
// 每个目录只能有一个包

func (node Node) Print() {
	fmt.Println(node.value)
}

func (node *Node) setValue(value int) {
	node.value = value
}

func TreeDemo() {
	var root Node
	root = Node{value: 3}
	root.left = &Node{}
	root.left.setValue(1)
	root.left.right = createNode(2)
	root.right = &Node{5, nil, nil}
	root.right.left = new(Node)
	nodes := []Node{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.Traverse()
}
