package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 上面的写法和下面的写法一样就是定义一个普通的方法,只不过在调用时的写法不一样,上面的是go语言特有的方法定义,有一个(值)接受者

func println(node treeNode) {
	fmt.Println(node.value)
}

// go语言中的接收者种类 值接收者 和 指针接收者
// go语言中只有值传递,意思是,无论是值传递还是指针传递都是复制了一份传递过去,值直接就是变量的字面值,而指针则是复制了指针中存放的内存地址
//
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node, Ignored")
		return
	}
	node.value = value
}

// 遍历树
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 工厂函数
func createNode(value int) *treeNode {
	// 在c++ 中return 一个函数中,这个是局部变量地址,会导致程序崩溃,在go中局部变量地址也是可以使用的
	return &treeNode{value: value}
}

// go语言中不需要构造函数,因为看下面的的代码提供了如此多的构造方法
// 有时候我们需要控制构造,可以添加工厂函数
func main() {
	root := treeNode{value: 3}

	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.left.right = createNode(8)
	// 两种不同定义调用的不同,
	root.print()
	fmt.Println()
	// 修改为指针接收者就可以修改原始struct中的数据,因为指针接收者将指针中存放的地址副本复制了一份,给了函数中的node但两份地址同时指向了
	// 原始的struct
	root.right.left.setValue(999)
	root.right.left.print() // 999
	fmt.Println()
	root.traverse()
}
