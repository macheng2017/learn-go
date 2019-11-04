package main

import (
	"fmt"
	"learngo/tree"
)

// 采用组合这种方式扩展类的功能
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

//postOrder()方法的接收者是一个指针接收者
//
//基于指针类型的方法，需要先声明一个指针变量，接收指针类型的值，编译器才能拿到具体的地址，调用其方法。

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

	fmt.Println()
	node := myTreeNode{&root}
	node.postOrder()
	fmt.Println()

}
