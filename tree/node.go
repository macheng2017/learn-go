package tree

import "fmt"

// 定义结构体(对象)
type Node struct {
	Value       int
	Left, Right *Node
}

// 工厂函数(构造函数)
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// define the method of struct
// root.print()
func (node Node) Print() {
	fmt.Println(node.Value)
}

//  等价于上面的方法,只是调用的方式不一样 print(root)
func Print(node Node) {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node. Ignored")
		return
	}
	node.Value = value
}
