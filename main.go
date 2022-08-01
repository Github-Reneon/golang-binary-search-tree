package main

import "fmt"

func AddNode(tree *Node, value int) *Node {
	if tree == nil {
		return &Node{
			Key: value,
		}
	}

	if value < tree.Key {
		tree.LeftNode  = AddNode(tree.LeftNode, value)
	} else if value > {

	}
}

func main() {
	node := AddNode(nil, 1)

	fmt.Println(node.Key)
}
