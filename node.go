package main

import (
	"fmt"
)

type Node struct {
	Key       int
	Name      string
	LeftNode  *Node
	RightNode *Node
}

type BST struct {
	root *Node
}

func (tree *BST) AddNode(value int, name string) *BST {
	if tree.root == nil {
		tree.root = &Node{Key: value, Name: name, LeftNode: nil, RightNode: nil}
	} else {
		tree.root.AddNode(value, name)
	}

	return tree
}

func (node *Node) AddNode(value int, name string) {
	if node == nil {
		return
	} else if value == node.Key {
		node.Name = name
	} else if value < node.Key {
		if node.LeftNode == nil {
			node.LeftNode = &Node{Key: value, Name: name, LeftNode: nil, RightNode: nil}
		} else {
			node.LeftNode.AddNode(value, name)
		}
	} else {
		if node.RightNode == nil {
			node.RightNode = &Node{Key: value, Name: name, LeftNode: nil, RightNode: nil}
		} else {
			node.RightNode.AddNode(value, name)
		}
	}
}

func (tree *BST) PrintInOrder() {
	if tree.root != nil {
		tree.root.PrintInOrder()
	}
}

func (node *Node) PrintInOrder() {
	if node == nil {
		return
	}
	node.LeftNode.PrintInOrder()
	print := fmt.Sprintf("%d, ", node.Key)
	fmt.Print(print)
	node.RightNode.PrintInOrder()
}

func (tree *BST) GetKey(name string) (int, int) {
	if tree.root == nil {
		return 0, -1
	}

	return tree.root.GetKey(name)
}

func (node *Node) GetKey(name string) (int, int) {
	if node == nil {
		return 0, -1
	}

	if node.Name == name {
		return node.Key, 0
	} else {
		key, err := node.LeftNode.GetKey(name)
		if err == 0 {
			return key, err
		}
		key, err = node.RightNode.GetKey(name)
		if err == 0 {
			return key, err
		}
		return 0, -1
	}
}

func (tree *BST) GetName(key int) (string, int) {
	if tree.root == nil {
		return "", -1
	}
	return tree.root.GetName(key)
}

func (node *Node) GetName(key int) (string, int) {
	if node == nil {
		return "", -1
	}

	if node.Key == key {
		return node.Name, 0
	} else {
		name, err := node.LeftNode.GetName(key)
		if err == 0 {
			return name, err
		}
		name, err = node.RightNode.GetName(key)
		if err == 0 {
			return name, err
		}
		return "", -1
	}
}