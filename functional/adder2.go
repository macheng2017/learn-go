package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(i int) int {
		// 函数体内有局部变量 i
		// sum不是这个函数体内定义的变量,而是环境的变量(外面定义的),在这里叫做自由变量
		// 通过这个sum可以连接到外部,如果sum是一个struct就可以网状形式连接到外部,
		//当所有变量都被连接组成一个网状的结构,就叫做闭包
		// 当函数内部引用了sum,函数会保存对sum的引用
		sum += i
		return sum
	}
}

func main() {
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... +%d = %d\n", i, a(i))
	}
}
