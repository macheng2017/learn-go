package main

import (
	"fmt"
	"learngo/tree"
)

func main() {
	root := tree.Node{}

	// left, right 是一个指针,需要接收一个地址
	root.Left = &tree.Node{Value: 0}

	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	// new func return pointer of the value
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	// defined slice
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	fmt.Println("nil pointer")
	var pRoot *tree.Node
	pRoot.SetValue(200)

	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	root.Traverse()

}
