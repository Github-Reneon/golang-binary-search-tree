package main

import (
	"fmt"
)

func main() {

	tree := &BST{}

	tree.AddNode(0, "root")

	tree.AddNode(30, "John")
	tree.AddNode(30, "James")
	tree.AddNode(60, "Henry")

	key, err := tree.GetKey("James")

	if err != -1 {
		fmt.Println(key)
	}

	name, err2 := tree.GetName(60)

	if err2 != -1 {
		fmt.Println(name)
	}

	tree.PrintInOrder()

}
